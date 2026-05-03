package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/plexusone/dark-factory-project-spec/types"
	"github.com/spf13/cobra"
)

var checkpointCmd = &cobra.Command{
	Use:   "checkpoint",
	Short: "Manage execution checkpoints",
	Long: `Manage checkpoints during spec execution.

Subcommands:
  list      List checkpoints for a feature
  show      Show checkpoint details
  approve   Approve a checkpoint
  reject    Reject a checkpoint

Examples:
  dfspec checkpoint list user-authentication
  dfspec checkpoint show user-authentication CP-001
  dfspec checkpoint approve user-authentication CP-001`,
}

var checkpointListCmd = &cobra.Command{
	Use:   "list <feature>",
	Short: "List checkpoints",
	Args:  cobra.ExactArgs(1),
	RunE:  runCheckpointList,
}

var checkpointShowCmd = &cobra.Command{
	Use:   "show <feature> <checkpoint-id>",
	Short: "Show checkpoint details",
	Args:  cobra.ExactArgs(2),
	RunE:  runCheckpointShow,
}

var checkpointApproveCmd = &cobra.Command{
	Use:   "approve <feature> <checkpoint-id>",
	Short: "Approve a checkpoint",
	Args:  cobra.ExactArgs(2),
	RunE:  runCheckpointApprove,
}

var checkpointRejectCmd = &cobra.Command{
	Use:   "reject <feature> <checkpoint-id>",
	Short: "Reject a checkpoint",
	Args:  cobra.ExactArgs(2),
	RunE:  runCheckpointReject,
}

var (
	checkpointNotes     string
	checkpointApprover  string
	checkpointRejectMsg string
)

func init() {
	checkpointCmd.AddCommand(checkpointListCmd)
	checkpointCmd.AddCommand(checkpointShowCmd)
	checkpointCmd.AddCommand(checkpointApproveCmd)
	checkpointCmd.AddCommand(checkpointRejectCmd)

	checkpointApproveCmd.Flags().StringVarP(&checkpointNotes, "notes", "n", "", "Approval notes")
	checkpointApproveCmd.Flags().StringVarP(&checkpointApprover, "approver", "a", "", "Approver name")

	checkpointRejectCmd.Flags().StringVarP(&checkpointRejectMsg, "message", "m", "", "Rejection reason (required)")
}

func loadPRD(feature string) (*types.PRD, error) {
	featureDir, err := findFeatureDir(feature)
	if err != nil {
		return nil, err
	}

	prdPath := filepath.Join(featureDir, "prd.json")
	data, err := os.ReadFile(prdPath)
	if err != nil {
		return nil, fmt.Errorf("PRD not found: %s", prdPath)
	}

	var prd types.PRD
	if err := json.Unmarshal(data, &prd); err != nil {
		return nil, fmt.Errorf("failed to parse PRD: %w", err)
	}

	return &prd, nil
}

func runCheckpointList(cmd *cobra.Command, args []string) error {
	feature := args[0]

	prd, err := loadPRD(feature)
	if err != nil {
		return err
	}

	state, _, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	if len(prd.Checkpoints) == 0 {
		printf("No checkpoints defined for '%s'\n", feature)
		return nil
	}

	printf("Checkpoints for %s:\n\n", feature)

	for _, cp := range prd.Checkpoints {
		status := "pending"
		if state != nil {
			for _, cs := range state.CheckpointStates {
				if cs.CheckpointID == cp.ID {
					status = string(cs.Status)
					break
				}
			}
		}

		statusIcon := "○"
		switch status {
		case "reached":
			statusIcon = "◉"
		case "passed":
			statusIcon = "✓"
		case "failed":
			statusIcon = "✗"
		case "blocked":
			statusIcon = "⚡"
		}

		printf("%s %s: %s\n", statusIcon, cp.ID, cp.Name)
		printf("    Type: %s\n", cp.Validation)
		printf("    After: %v\n", cp.AfterRequirements)
		printf("    Status: %s\n", status)
		printf("\n")
	}

	return nil
}

func runCheckpointShow(cmd *cobra.Command, args []string) error {
	feature := args[0]
	checkpointID := args[1]

	prd, err := loadPRD(feature)
	if err != nil {
		return err
	}

	// Find checkpoint in PRD
	var checkpoint *types.Checkpoint
	for i := range prd.Checkpoints {
		if prd.Checkpoints[i].ID == checkpointID {
			checkpoint = &prd.Checkpoints[i]
			break
		}
	}

	if checkpoint == nil {
		return fmt.Errorf("checkpoint '%s' not found in PRD", checkpointID)
	}

	state, _, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	printf("Checkpoint: %s\n", checkpoint.ID)
	printf("Name: %s\n", checkpoint.Name)
	printf("Description: %s\n", checkpoint.Description)
	printf("Validation: %s\n", checkpoint.Validation)
	printf("Pause on Discovery: %t\n", checkpoint.PauseOnDiscovery)
	printf("\nAfter Requirements:\n")
	for _, reqID := range checkpoint.AfterRequirements {
		reqStatus := "pending"
		if state != nil {
			for _, rs := range state.RequirementStates {
				if rs.RequirementID == reqID {
					reqStatus = string(rs.Status)
					break
				}
			}
		}
		printf("  - %s (%s)\n", reqID, reqStatus)
	}

	if len(checkpoint.ValidationCriteria) > 0 {
		printf("\nValidation Criteria:\n")
		for _, crit := range checkpoint.ValidationCriteria {
			printf("  - %s\n", crit)
		}
	}

	// Show execution state if exists
	if state != nil {
		for _, cs := range state.CheckpointStates {
			if cs.CheckpointID == checkpointID {
				printf("\nExecution State:\n")
				printf("  Status: %s\n", cs.Status)
				if cs.ReachedAt != "" {
					printf("  Reached: %s\n", cs.ReachedAt)
				}
				if cs.ResolvedAt != "" {
					printf("  Resolved: %s\n", cs.ResolvedAt)
				}
				if cs.WaitingFor != "" {
					printf("  Waiting for: %s\n", cs.WaitingFor)
				}
				if cs.Result != nil {
					printf("  Passed: %t\n", cs.Result.Passed)
					if cs.Result.Notes != "" {
						printf("  Notes: %s\n", cs.Result.Notes)
					}
				}
				break
			}
		}
	}

	return nil
}

func runCheckpointApprove(cmd *cobra.Command, args []string) error {
	feature := args[0]
	checkpointID := args[1]

	state, execPath, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	if state == nil {
		return fmt.Errorf("no execution state found for '%s'", feature)
	}

	// Find checkpoint state
	var checkpointState *types.CheckpointState
	for i := range state.CheckpointStates {
		if state.CheckpointStates[i].CheckpointID == checkpointID {
			checkpointState = &state.CheckpointStates[i]
			break
		}
	}

	if checkpointState == nil {
		return fmt.Errorf("checkpoint '%s' not found in execution state", checkpointID)
	}

	if checkpointState.Status != types.CheckpointReached {
		return fmt.Errorf("checkpoint '%s' is not in 'reached' status (current: %s)", checkpointID, checkpointState.Status)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	approver := checkpointApprover
	if approver == "" {
		approver = "CLI user"
	}

	checkpointState.Status = types.CheckpointPassed
	checkpointState.ResolvedAt = now
	checkpointState.Result = &types.CheckpointResult{
		ValidatedAt: now,
		ValidatedBy: approver,
		Passed:      true,
		Notes:       checkpointNotes,
	}

	// Update execution state
	if state.CurrentCheckpoint == checkpointID {
		state.CurrentCheckpoint = ""
		state.PausedReason = ""
		state.ResumeInstructions = ""
	}

	// If execution was paused at this checkpoint, mark as in_progress
	if state.Status == types.ExecutionPaused {
		state.Status = types.ExecutionInProgress
	}

	// Add log entry
	state.ExecutionLog = append(state.ExecutionLog, types.ExecutionLogEntry{
		Timestamp:    now,
		Event:        types.EventCheckpointPassed,
		Details:      fmt.Sprintf("Checkpoint approved by %s", approver),
		CheckpointID: checkpointID,
	})

	if err := saveExecutionState(state, execPath); err != nil {
		return err
	}

	printf("✓ Checkpoint '%s' approved\n", checkpointID)
	if checkpointNotes != "" {
		printf("  Notes: %s\n", checkpointNotes)
	}
	printf("\nExecution can continue. Use: /dfspec-exec %s --resume\n", feature)

	return nil
}

func runCheckpointReject(cmd *cobra.Command, args []string) error {
	feature := args[0]
	checkpointID := args[1]

	if checkpointRejectMsg == "" {
		return fmt.Errorf("--message is required when rejecting a checkpoint")
	}

	state, execPath, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	if state == nil {
		return fmt.Errorf("no execution state found for '%s'", feature)
	}

	// Find checkpoint state
	var checkpointState *types.CheckpointState
	for i := range state.CheckpointStates {
		if state.CheckpointStates[i].CheckpointID == checkpointID {
			checkpointState = &state.CheckpointStates[i]
			break
		}
	}

	if checkpointState == nil {
		return fmt.Errorf("checkpoint '%s' not found in execution state", checkpointID)
	}

	if checkpointState.Status != types.CheckpointReached {
		return fmt.Errorf("checkpoint '%s' is not in 'reached' status (current: %s)", checkpointID, checkpointState.Status)
	}

	now := time.Now().UTC().Format(time.RFC3339)

	checkpointState.Status = types.CheckpointFailed
	checkpointState.ResolvedAt = now
	checkpointState.Result = &types.CheckpointResult{
		ValidatedAt:    now,
		ValidatedBy:    "CLI user",
		Passed:         false,
		Notes:          checkpointRejectMsg,
		BlockingIssues: []string{checkpointRejectMsg},
	}

	// Update execution state
	state.Status = types.ExecutionFailed
	state.PausedReason = fmt.Sprintf("Checkpoint %s rejected: %s", checkpointID, checkpointRejectMsg)

	// Add log entry
	state.ExecutionLog = append(state.ExecutionLog, types.ExecutionLogEntry{
		Timestamp:    now,
		Event:        types.EventCheckpointFailed,
		Details:      fmt.Sprintf("Checkpoint rejected: %s", checkpointRejectMsg),
		CheckpointID: checkpointID,
	})

	if err := saveExecutionState(state, execPath); err != nil {
		return err
	}

	printf("✗ Checkpoint '%s' rejected\n", checkpointID)
	printf("  Reason: %s\n", checkpointRejectMsg)
	printf("\nExecution stopped.\n")

	return nil
}
