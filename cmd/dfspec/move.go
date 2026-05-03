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

var moveCmd = &cobra.Command{
	Use:   "move <feature> <status>",
	Short: "Move a feature to a different status",
	Long: `Move a feature between status directories.

Valid statuses: active, backlog, completed, archived

Examples:
  dfspec move query-caching active      # Start working on feature
  dfspec move query-caching completed   # Mark feature as done
  dfspec move old-feature archived      # Archive abandoned feature`,
	Args: cobra.ExactArgs(2),
	RunE: runMove,
}

var (
	moveReason    string
	moveVersion   string
	moveSupersede string
)

func init() {
	moveCmd.Flags().StringVarP(&moveReason, "reason", "r", "", "Reason for archiving (required for archived)")
	moveCmd.Flags().StringVarP(&moveVersion, "version", "v", "", "Release version (for completed)")
	moveCmd.Flags().StringVar(&moveSupersede, "superseded-by", "", "Feature that supersedes this one (for archived)")
}

func runMove(cmd *cobra.Command, args []string) error {
	feature := args[0]
	targetStatus := args[1]

	// Validate target status
	validStatuses := map[string]bool{"active": true, "backlog": true, "completed": true, "archived": true}
	if !validStatuses[targetStatus] {
		return fmt.Errorf("invalid status '%s', must be one of: active, backlog, completed, archived", targetStatus)
	}

	// Require reason for archiving
	if targetStatus == "archived" && moveReason == "" {
		return fmt.Errorf("--reason is required when archiving a feature")
	}

	// Find current location
	var currentStatus string
	var currentPath string
	for _, status := range []string{"active", "backlog", "completed", "archived"} {
		checkPath := filepath.Join(specsDir, status, feature)
		if _, err := os.Stat(checkPath); err == nil {
			currentStatus = status
			currentPath = checkPath
			break
		}
	}

	if currentStatus == "" {
		return fmt.Errorf("feature '%s' not found in any status directory", feature)
	}

	if currentStatus == targetStatus {
		return fmt.Errorf("feature '%s' is already in %s/", feature, targetStatus)
	}

	// Move directory
	targetPath := filepath.Join(specsDir, targetStatus, feature)

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	if err := os.Rename(currentPath, targetPath); err != nil {
		return fmt.Errorf("failed to move directory: %w", err)
	}

	// Update roadmap
	if err := updateRoadmapMove(feature, currentStatus, targetStatus); err != nil {
		errorf("Warning: failed to update roadmap.json: %v\n", err)
	}

	printf("Moved '%s' from %s/ to %s/\n", feature, currentStatus, targetStatus)

	// Print next steps based on transition
	switch targetStatus {
	case "active":
		printf("\nFeature is now active. Next steps:\n")
		printf("  1. Create/update PRD: %s/prd.json\n", targetPath)
		printf("  2. Implement the feature\n")
		printf("  3. Run 'dfspec move %s completed' when done\n", feature)
	case "completed":
		printf("\nFeature marked as completed.\n")
		if moveVersion != "" {
			printf("  Released in: %s\n", moveVersion)
		}
	case "archived":
		printf("\nFeature archived.\n")
		printf("  Reason: %s\n", moveReason)
		if moveSupersede != "" {
			printf("  Superseded by: %s\n", moveSupersede)
		}
	}

	return nil
}

func updateRoadmapMove(feature, fromStatus, toStatus string) error {
	roadmapPath := filepath.Join(specsDir, "roadmap.json")

	data, err := os.ReadFile(roadmapPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No roadmap to update
		}
		return err
	}

	var roadmap types.Roadmap
	if err := json.Unmarshal(data, &roadmap); err != nil {
		return err
	}

	roadmap.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	now := time.Now().UTC().Format(time.RFC3339)

	// Remove from source list
	switch fromStatus {
	case "active":
		roadmap.Active = removeActive(roadmap.Active, feature)
	case "backlog":
		roadmap.Backlog = removeBacklog(roadmap.Backlog, feature)
	case "completed":
		roadmap.Completed = removeCompleted(roadmap.Completed, feature)
	case "archived":
		roadmap.Archived = removeArchived(roadmap.Archived, feature)
	}

	// Add to target list
	switch toStatus {
	case "active":
		maxPriority := 0
		for _, f := range roadmap.Active {
			if f.Priority > maxPriority {
				maxPriority = f.Priority
			}
		}
		roadmap.Active = append(roadmap.Active, types.ActiveFeature{
			Feature:   feature,
			Priority:  maxPriority + 1,
			StartedAt: now,
		})
	case "backlog":
		maxPriority := 0
		for _, f := range roadmap.Backlog {
			if f.Priority > maxPriority {
				maxPriority = f.Priority
			}
		}
		roadmap.Backlog = append(roadmap.Backlog, types.BacklogFeature{
			Feature:  feature,
			Priority: maxPriority + 1,
		})
	case "completed":
		roadmap.Completed = append(roadmap.Completed, types.CompletedFeature{
			Feature:     feature,
			CompletedAt: now,
			ReleasedIn:  moveVersion,
		})
	case "archived":
		roadmap.Archived = append(roadmap.Archived, types.ArchivedFeature{
			Feature:      feature,
			ArchivedAt:   now,
			Reason:       moveReason,
			SupersededBy: moveSupersede,
		})
	}

	newData, err := json.MarshalIndent(roadmap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(roadmapPath, newData, 0644)
}

func removeActive(list []types.ActiveFeature, feature string) []types.ActiveFeature {
	result := make([]types.ActiveFeature, 0, len(list))
	for _, f := range list {
		if f.Feature != feature {
			result = append(result, f)
		}
	}
	return result
}

func removeBacklog(list []types.BacklogFeature, feature string) []types.BacklogFeature {
	result := make([]types.BacklogFeature, 0, len(list))
	for _, f := range list {
		if f.Feature != feature {
			result = append(result, f)
		}
	}
	return result
}

func removeCompleted(list []types.CompletedFeature, feature string) []types.CompletedFeature {
	result := make([]types.CompletedFeature, 0, len(list))
	for _, f := range list {
		if f.Feature != feature {
			result = append(result, f)
		}
	}
	return result
}

func removeArchived(list []types.ArchivedFeature, feature string) []types.ArchivedFeature {
	result := make([]types.ArchivedFeature, 0, len(list))
	for _, f := range list {
		if f.Feature != feature {
			result = append(result, f)
		}
	}
	return result
}
