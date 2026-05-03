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

var executeCmd = &cobra.Command{
	Use:   "execute <feature>",
	Short: "Manage spec execution",
	Long: `Manage the execution of a PRD specification.

Subcommands:
  status    Show execution status
  log       Show execution log
  start     Start execution (creates execution.json)
  pause     Pause execution
  resume    Resume paused execution

Examples:
  dfspec execute status user-authentication
  dfspec execute log user-authentication
  dfspec execute start user-authentication`,
}

var executeStatusCmd = &cobra.Command{
	Use:   "status <feature>",
	Short: "Show execution status",
	Args:  cobra.ExactArgs(1),
	RunE:  runExecuteStatus,
}

var executeLogCmd = &cobra.Command{
	Use:   "log <feature>",
	Short: "Show execution log",
	Args:  cobra.ExactArgs(1),
	RunE:  runExecuteLog,
}

var executeStartCmd = &cobra.Command{
	Use:   "start <feature>",
	Short: "Initialize execution state",
	Args:  cobra.ExactArgs(1),
	RunE:  runExecuteStart,
}

var executePauseCmd = &cobra.Command{
	Use:   "pause <feature>",
	Short: "Pause execution",
	Args:  cobra.ExactArgs(1),
	RunE:  runExecutePause,
}

var (
	executeLogLimit int
	pauseReason     string
)

func init() {
	executeCmd.AddCommand(executeStatusCmd)
	executeCmd.AddCommand(executeLogCmd)
	executeCmd.AddCommand(executeStartCmd)
	executeCmd.AddCommand(executePauseCmd)

	executeLogCmd.Flags().IntVarP(&executeLogLimit, "limit", "n", 10, "Number of log entries to show")
	executePauseCmd.Flags().StringVarP(&pauseReason, "reason", "r", "", "Reason for pausing")
}

func findFeatureDir(feature string) (string, error) {
	for _, status := range []string{"active", "backlog", "completed", "archived"} {
		path := filepath.Join(specsDir, status, feature)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("feature '%s' not found", feature)
}

func loadExecutionState(feature string) (*types.ExecutionState, string, error) {
	featureDir, err := findFeatureDir(feature)
	if err != nil {
		return nil, "", err
	}

	execPath := filepath.Join(featureDir, "execution.json")
	data, err := os.ReadFile(execPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, execPath, nil // No execution state yet
		}
		return nil, "", err
	}

	var state types.ExecutionState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, "", err
	}

	return &state, execPath, nil
}

func saveExecutionState(state *types.ExecutionState, path string) error {
	state.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func runExecuteStatus(cmd *cobra.Command, args []string) error {
	feature := args[0]

	state, _, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	if state == nil {
		printf("%s: NOT STARTED\n", feature)
		printf("  No execution state found.\n")
		printf("  Start with: dfspec execute start %s\n", feature)
		return nil
	}

	// Calculate progress
	implemented := 0
	blocked := 0
	pending := 0
	for _, rs := range state.RequirementStates {
		switch rs.Status {
		case types.StatusImplemented, types.StatusValidated:
			implemented++
		case types.StatusBlocked:
			blocked++
		default:
			pending++
		}
	}
	total := len(state.RequirementStates)

	printf("%s: %s\n", feature, state.Status)
	printf("  Progress: %d/%d requirements (%.0f%%)\n",
		implemented, total, float64(implemented)/float64(total)*100)

	if state.CurrentCheckpoint != "" {
		printf("  Checkpoint: %s\n", state.CurrentCheckpoint)
	}

	if blocked > 0 {
		printf("  Blocked: %d requirements\n", blocked)
	}

	if len(state.Discoveries) > 0 {
		unresolvedCount := 0
		for _, d := range state.Discoveries {
			if d.Resolution == nil {
				unresolvedCount++
			}
		}
		if unresolvedCount > 0 {
			printf("  Discoveries: %d pending\n", unresolvedCount)
		}
	}

	if state.PausedReason != "" {
		printf("  Paused: %s\n", state.PausedReason)
	}

	if state.ResumeInstructions != "" {
		printf("  Resume: %s\n", state.ResumeInstructions)
	}

	return nil
}

func runExecuteLog(cmd *cobra.Command, args []string) error {
	feature := args[0]

	state, _, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	if state == nil {
		return fmt.Errorf("no execution state found for '%s'", feature)
	}

	if len(state.ExecutionLog) == 0 {
		printf("No execution log entries.\n")
		return nil
	}

	// Show most recent entries
	start := 0
	if executeLogLimit > 0 && len(state.ExecutionLog) > executeLogLimit {
		start = len(state.ExecutionLog) - executeLogLimit
	}

	printf("Execution log for %s:\n\n", feature)
	for _, entry := range state.ExecutionLog[start:] {
		printf("%s [%s] %s\n", entry.Timestamp[:19], entry.Event, entry.Details)
		if entry.RequirementID != "" {
			printf("  Requirement: %s\n", entry.RequirementID)
		}
		if entry.CheckpointID != "" {
			printf("  Checkpoint: %s\n", entry.CheckpointID)
		}
	}

	return nil
}

func runExecuteStart(cmd *cobra.Command, args []string) error {
	feature := args[0]

	featureDir, err := findFeatureDir(feature)
	if err != nil {
		return err
	}

	// Check if PRD exists
	prdPath := filepath.Join(featureDir, "prd.json")
	prdData, err := os.ReadFile(prdPath)
	if err != nil {
		return fmt.Errorf("PRD not found: %s", prdPath)
	}

	var prd types.PRD
	if err := json.Unmarshal(prdData, &prd); err != nil {
		return fmt.Errorf("failed to parse PRD: %w", err)
	}

	// Check if already started
	execPath := filepath.Join(featureDir, "execution.json")
	if _, err := os.Stat(execPath); err == nil {
		return fmt.Errorf("execution already started; use 'dfspec execute status %s' to check status", feature)
	}

	// Create initial execution state
	now := time.Now().UTC().Format(time.RFC3339)

	state := types.ExecutionState{
		PRDReference: prd.Metadata.ID,
		Feature:      feature,
		StartedAt:    now,
		UpdatedAt:    now,
		Status:       types.ExecutionNotStarted,
	}

	// Initialize requirement states
	for _, fr := range prd.FunctionalRequirements {
		state.RequirementStates = append(state.RequirementStates, types.RequirementState{
			RequirementID: fr.ID,
			Status:        types.StatusPending,
		})
	}

	// Initialize checkpoint states
	for _, cp := range prd.Checkpoints {
		state.CheckpointStates = append(state.CheckpointStates, types.CheckpointState{
			CheckpointID: cp.ID,
			Status:       types.CheckpointPending,
		})
	}

	// Add log entry
	state.ExecutionLog = append(state.ExecutionLog, types.ExecutionLogEntry{
		Timestamp: now,
		Event:     types.EventStarted,
		Details:   fmt.Sprintf("Execution initialized for %s", prd.Metadata.Title),
	})

	// Save state
	if err := saveExecutionState(&state, execPath); err != nil {
		return err
	}

	printf("Execution initialized for '%s'\n", feature)
	printf("  PRD: %s (%s)\n", prd.Metadata.Title, prd.Metadata.ID)
	printf("  Requirements: %d\n", len(prd.FunctionalRequirements))
	printf("  Checkpoints: %d\n", len(prd.Checkpoints))
	printf("\nTo execute, use: /dfspec-exec %s\n", feature)

	return nil
}

func runExecutePause(cmd *cobra.Command, args []string) error {
	feature := args[0]

	state, execPath, err := loadExecutionState(feature)
	if err != nil {
		return err
	}

	if state == nil {
		return fmt.Errorf("no execution state found for '%s'", feature)
	}

	if state.Status == types.ExecutionPaused {
		return fmt.Errorf("execution is already paused")
	}

	if state.Status == types.ExecutionCompleted {
		return fmt.Errorf("execution is already completed")
	}

	now := time.Now().UTC().Format(time.RFC3339)

	state.Status = types.ExecutionPaused
	state.PausedReason = pauseReason
	if pauseReason == "" {
		state.PausedReason = "Manually paused"
	}
	state.ResumeInstructions = fmt.Sprintf("/dfspec-exec %s --resume", feature)

	state.ExecutionLog = append(state.ExecutionLog, types.ExecutionLogEntry{
		Timestamp: now,
		Event:     types.EventPaused,
		Details:   state.PausedReason,
	})

	if err := saveExecutionState(state, execPath); err != nil {
		return err
	}

	printf("Execution paused for '%s'\n", feature)
	printf("  Reason: %s\n", state.PausedReason)
	printf("  Resume: /dfspec-exec %s --resume\n", feature)

	return nil
}
