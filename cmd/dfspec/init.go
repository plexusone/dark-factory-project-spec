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

var initCmd = &cobra.Command{
	Use:   "init <feature-name>",
	Short: "Initialize a new feature",
	Long: `Create a new feature directory with scaffolded MRD.

The feature is created in specs/backlog/ by default.
Use --status to create in a different status directory.

Example:
  dfspec init query-caching
  dfspec init user-auth --status active`,
	Args: cobra.ExactArgs(1),
	RunE: runInit,
}

var (
	initStatus string
	initTitle  string
	initAuthor string
)

func init() {
	initCmd.Flags().StringVarP(&initStatus, "status", "s", "backlog", "Initial status (active, backlog)")
	initCmd.Flags().StringVarP(&initTitle, "title", "t", "", "Feature title (defaults to feature name)")
	initCmd.Flags().StringVarP(&initAuthor, "author", "a", "", "Author name")
}

func runInit(cmd *cobra.Command, args []string) error {
	feature := args[0]

	// Validate feature name
	if !isKebabCase(feature) {
		return fmt.Errorf("feature name must be kebab-case (e.g., 'query-caching'), got '%s'", feature)
	}

	// Validate status
	if initStatus != "active" && initStatus != "backlog" {
		return fmt.Errorf("status must be 'active' or 'backlog', got '%s'", initStatus)
	}

	// Check if feature already exists anywhere
	for _, status := range []string{"active", "backlog", "completed", "archived"} {
		checkPath := filepath.Join(specsDir, status, feature)
		if _, err := os.Stat(checkPath); err == nil {
			return fmt.Errorf("feature '%s' already exists in %s/", feature, status)
		}
	}

	// Create feature directory
	featureDir := filepath.Join(specsDir, initStatus, feature)
	if err := os.MkdirAll(featureDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Generate MRD
	title := initTitle
	if title == "" {
		title = feature // Will be converted to title case by user
	}

	now := time.Now().UTC().Format(time.RFC3339)
	year := time.Now().Year()

	mrd := types.MRD{
		Metadata: types.DocumentMetadata{
			ID:        fmt.Sprintf("MRD-%d-%s", year, feature),
			Title:     title,
			Feature:   feature,
			Version:   "0.1.0",
			Status:    types.StatusDraft,
			Authors:   []string{initAuthor},
			CreatedAt: now,
			UpdatedAt: now,
		},
		ProblemStatement: types.ProblemStatement{
			Summary:      "TODO: Describe the problem",
			CurrentState: "TODO: Describe how things work today",
			DesiredState: "TODO: Describe how things should work",
			Impact:       "TODO: Describe the business impact",
		},
		TargetUsers: []types.TargetUser{
			{
				Name:        "TODO: Primary User",
				Description: "TODO: Who they are",
				Needs:       []string{"TODO: What they need"},
				Priority:    types.PriorityPrimary,
			},
		},
		SuccessMetrics: []types.SuccessMetric{
			{
				ID:                "SM-001",
				Name:              "TODO: Metric Name",
				Description:       "TODO: What is being measured",
				TargetValue:       "TODO: Target",
				MeasurementMethod: "TODO: How to measure",
			},
		},
		HighLevelRequirements: []types.HighLevelRequirement{
			{
				ID:          "HLR-001",
				Description: "TODO: High-level requirement",
				Rationale:   "TODO: Why this is needed",
				Priority:    types.PriorityMustHave,
			},
		},
		NonGoals: []string{
			"TODO: What is explicitly out of scope",
		},
	}

	// Filter empty author
	if initAuthor == "" {
		mrd.Metadata.Authors = []string{}
	}

	// Write MRD
	mrdPath := filepath.Join(featureDir, "mrd.json")
	mrdData, err := json.MarshalIndent(mrd, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal MRD: %w", err)
	}
	if err := os.WriteFile(mrdPath, mrdData, 0644); err != nil {
		return fmt.Errorf("failed to write MRD: %w", err)
	}

	// Update roadmap
	if err := addToRoadmap(feature, initStatus); err != nil {
		// Don't fail, just warn
		errorf("Warning: failed to update roadmap.json: %v\n", err)
	}

	// Create docs directory
	docsFeatureDir := filepath.Join(docsDir, "design", feature)
	if err := os.MkdirAll(docsFeatureDir, 0755); err != nil {
		errorf("Warning: failed to create docs directory: %v\n", err)
	}

	printf("Created feature '%s' in %s/\n", feature, initStatus)
	printf("\nFiles created:\n")
	printf("  %s\n", mrdPath)
	printf("  %s/\n", docsFeatureDir)
	printf("\nNext steps:\n")
	printf("  1. Edit %s to define requirements\n", mrdPath)
	printf("  2. Run 'dfspec validate' to check validity\n")
	printf("  3. When ready, run 'dfspec move %s active' to start work\n", feature)

	return nil
}

func addToRoadmap(feature string, status string) error {
	roadmapPath := filepath.Join(specsDir, "roadmap.json")

	// Read existing roadmap or create new
	var roadmap types.Roadmap
	data, err := os.ReadFile(roadmapPath)
	if err != nil {
		if os.IsNotExist(err) {
			roadmap = types.Roadmap{
				Version:   "1.0.0",
				UpdatedAt: time.Now().UTC().Format(time.RFC3339),
				Active:    []types.ActiveFeature{},
				Backlog:   []types.BacklogFeature{},
				Completed: []types.CompletedFeature{},
				Archived:  []types.ArchivedFeature{},
			}
		} else {
			return err
		}
	} else {
		if err := json.Unmarshal(data, &roadmap); err != nil {
			return err
		}
	}

	// Add to appropriate list
	roadmap.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	switch status {
	case "active":
		// Find max priority
		maxPriority := 0
		for _, f := range roadmap.Active {
			if f.Priority > maxPriority {
				maxPriority = f.Priority
			}
		}
		roadmap.Active = append(roadmap.Active, types.ActiveFeature{
			Feature:   feature,
			Priority:  maxPriority + 1,
			StartedAt: time.Now().UTC().Format(time.RFC3339),
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
	}

	// Write back
	newData, err := json.MarshalIndent(roadmap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(roadmapPath, newData, 0644)
}
