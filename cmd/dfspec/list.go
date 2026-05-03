package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/plexusone/dark-factory-project-spec/types"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List features",
	Long: `List features by status.

Without arguments, shows all features grouped by status.
With a status argument, shows only features in that status.

Examples:
  dfspec list           # List all features
  dfspec list active    # List only active features`,
	Args: cobra.MaximumNArgs(1),
	RunE: runList,
}

var listJSON bool

func init() {
	listCmd.Flags().BoolVar(&listJSON, "json", false, "Output as JSON")
}

func runList(cmd *cobra.Command, args []string) error {
	// Try to read roadmap
	roadmapPath := filepath.Join(specsDir, "roadmap.json")
	data, err := os.ReadFile(roadmapPath)

	var roadmap *types.Roadmap
	if err == nil {
		roadmap = &types.Roadmap{}
		if err := json.Unmarshal(data, roadmap); err != nil {
			roadmap = nil
		}
	}

	// Filter by status if specified
	var filterStatus string
	if len(args) > 0 {
		filterStatus = args[0]
		validStatuses := map[string]bool{"active": true, "backlog": true, "completed": true, "archived": true}
		if !validStatuses[filterStatus] {
			return fmt.Errorf("invalid status '%s', must be one of: active, backlog, completed, archived", filterStatus)
		}
	}

	if listJSON {
		return listJSONOutput(roadmap, filterStatus)
	}

	return listTextOutput(roadmap, filterStatus)
}

func listTextOutput(roadmap *types.Roadmap, filterStatus string) error {
	statuses := []string{"active", "backlog", "completed", "archived"}
	if filterStatus != "" {
		statuses = []string{filterStatus}
	}

	for _, status := range statuses {
		features := getFeaturesForStatus(roadmap, status)

		if len(features) == 0 && filterStatus == "" {
			continue
		}

		printf("\n%s (%d):\n", statusLabel(status), len(features))

		if len(features) == 0 {
			printf("  (none)\n")
			continue
		}

		for _, f := range features {
			printf("  • %s\n", f.Name)
			if f.Priority > 0 {
				printf("    Priority: %d\n", f.Priority)
			}
			if f.Owner != "" {
				printf("    Owner: %s\n", f.Owner)
			}
			if f.Version != "" {
				printf("    Version: %s\n", f.Version)
			}
			if f.Note != "" {
				printf("    Note: %s\n", f.Note)
			}
		}
	}

	printf("\n")
	return nil
}

type featureInfo struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Priority int    `json:"priority,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Version  string `json:"version,omitempty"`
	Note     string `json:"note,omitempty"`
}

func getFeaturesForStatus(roadmap *types.Roadmap, status string) []featureInfo {
	var features []featureInfo

	// First try roadmap
	if roadmap != nil {
		switch status {
		case "active":
			for _, f := range roadmap.Active {
				features = append(features, featureInfo{
					Name:     f.Feature,
					Status:   status,
					Priority: f.Priority,
					Owner:    f.Owner,
					Version:  f.TargetVersion,
					Note:     f.Notes,
				})
			}
		case "backlog":
			for _, f := range roadmap.Backlog {
				features = append(features, featureInfo{
					Name:     f.Feature,
					Status:   status,
					Priority: f.Priority,
					Version:  f.EstimatedVersion,
					Note:     f.Notes,
				})
			}
		case "completed":
			for _, f := range roadmap.Completed {
				features = append(features, featureInfo{
					Name:    f.Feature,
					Status:  status,
					Version: f.ReleasedIn,
					Note:    f.Notes,
				})
			}
		case "archived":
			for _, f := range roadmap.Archived {
				features = append(features, featureInfo{
					Name:   f.Feature,
					Status: status,
					Note:   f.Reason,
				})
			}
		}
		return features
	}

	// Fall back to directory listing
	statusDir := filepath.Join(specsDir, status)
	entries, err := os.ReadDir(statusDir)
	if err != nil {
		return features
	}

	for _, entry := range entries {
		if entry.IsDir() {
			features = append(features, featureInfo{
				Name:   entry.Name(),
				Status: status,
			})
		}
	}

	return features
}

func listJSONOutput(roadmap *types.Roadmap, filterStatus string) error {
	statuses := []string{"active", "backlog", "completed", "archived"}
	if filterStatus != "" {
		statuses = []string{filterStatus}
	}

	result := make(map[string][]featureInfo)
	for _, status := range statuses {
		result[status] = getFeaturesForStatus(roadmap, status)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	printf("%s\n", data)
	return nil
}

func statusLabel(status string) string {
	labels := map[string]string{
		"active":    "Active",
		"backlog":   "Backlog",
		"completed": "Completed",
		"archived":  "Archived",
	}
	if label, ok := labels[status]; ok {
		return label
	}
	return status
}
