package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/plexusone/dark-factory-project-spec/types"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate <type>",
	Short: "Generate documentation",
	Long: `Generate human-readable documentation from specs.

Types:
  roadmap    Generate ROADMAP.md from roadmap.json
  index      Generate index.md files for feature directories`,
}

var generateRoadmapCmd = &cobra.Command{
	Use:   "roadmap",
	Short: "Generate ROADMAP.md from roadmap.json",
	RunE:  runGenerateRoadmap,
}

var generateIndexCmd = &cobra.Command{
	Use:   "index [feature]",
	Short: "Generate index.md for feature docs",
	Long: `Generate index.md files for feature documentation directories.

Without arguments, generates index.md for all features.
With a feature argument, generates only for that feature.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runGenerateIndex,
}

var (
	generateOutput string
)

func init() {
	generateCmd.AddCommand(generateRoadmapCmd)
	generateCmd.AddCommand(generateIndexCmd)

	generateRoadmapCmd.Flags().StringVarP(&generateOutput, "output", "o", "ROADMAP.md", "Output file path")
}

func runGenerateRoadmap(cmd *cobra.Command, args []string) error {
	roadmapPath := filepath.Join(specsDir, "roadmap.json")
	data, err := os.ReadFile(roadmapPath)
	if err != nil {
		return fmt.Errorf("failed to read roadmap.json: %w", err)
	}

	var roadmap types.Roadmap
	if err := json.Unmarshal(data, &roadmap); err != nil {
		return fmt.Errorf("failed to parse roadmap.json: %w", err)
	}

	var sb strings.Builder

	sb.WriteString("# Roadmap\n\n")
	sb.WriteString("_Generated from specs/roadmap.json_\n\n")

	// Active
	sb.WriteString("## Active\n\n")
	if len(roadmap.Active) == 0 {
		sb.WriteString("_No active features_\n\n")
	} else {
		sb.WriteString("| Priority | Feature | Owner | Target |\n")
		sb.WriteString("|----------|---------|-------|--------|\n")
		for _, f := range roadmap.Active {
			owner := f.Owner
			if owner == "" {
				owner = "-"
			}
			target := f.TargetVersion
			if target == "" {
				target = "-"
			}
			_, _ = fmt.Fprintf(&sb, "| %d | [%s](specs/active/%s/) | %s | %s |\n",
				f.Priority, f.Feature, f.Feature, owner, target)
		}
		sb.WriteString("\n")
	}

	// Backlog
	sb.WriteString("## Backlog\n\n")
	if len(roadmap.Backlog) == 0 {
		sb.WriteString("_No backlog features_\n\n")
	} else {
		sb.WriteString("| Priority | Feature | Blocked By | Est. Version |\n")
		sb.WriteString("|----------|---------|------------|---------------|\n")
		for _, f := range roadmap.Backlog {
			blockedBy := "-"
			if len(f.BlockedBy) > 0 {
				blockedBy = strings.Join(f.BlockedBy, ", ")
			}
			version := f.EstimatedVersion
			if version == "" {
				version = "-"
			}
			_, _ = fmt.Fprintf(&sb, "| %d | [%s](specs/backlog/%s/) | %s | %s |\n",
				f.Priority, f.Feature, f.Feature, blockedBy, version)
		}
		sb.WriteString("\n")
	}

	// Completed
	sb.WriteString("## Completed\n\n")
	if len(roadmap.Completed) == 0 {
		sb.WriteString("_No completed features_\n\n")
	} else {
		sb.WriteString("| Feature | Released In | Completed |\n")
		sb.WriteString("|---------|-------------|------------|\n")
		for _, f := range roadmap.Completed {
			version := f.ReleasedIn
			if version == "" {
				version = "-"
			}
			completed := f.CompletedAt
			if len(completed) > 10 {
				completed = completed[:10] // Just date
			}
			_, _ = fmt.Fprintf(&sb, "| [%s](specs/completed/%s/) | %s | %s |\n",
				f.Feature, f.Feature, version, completed)
		}
		sb.WriteString("\n")
	}

	// Archived
	if len(roadmap.Archived) > 0 {
		sb.WriteString("## Archived\n\n")
		sb.WriteString("| Feature | Reason | Superseded By |\n")
		sb.WriteString("|---------|--------|---------------|\n")
		for _, f := range roadmap.Archived {
			superseded := f.SupersededBy
			if superseded == "" {
				superseded = "-"
			}
			_, _ = fmt.Fprintf(&sb, "| %s | %s | %s |\n",
				f.Feature, f.Reason, superseded)
		}
		sb.WriteString("\n")
	}

	// Write output
	if err := os.WriteFile(generateOutput, []byte(sb.String()), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", generateOutput, err)
	}

	printf("Generated %s\n", generateOutput)
	return nil
}

func runGenerateIndex(cmd *cobra.Command, args []string) error {
	var features []string

	if len(args) > 0 {
		features = []string{args[0]}
	} else {
		// Find all features
		for _, status := range []string{"active", "backlog", "completed", "archived"} {
			statusDir := filepath.Join(specsDir, status)
			entries, err := os.ReadDir(statusDir)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				if entry.IsDir() {
					features = append(features, entry.Name())
				}
			}
		}
	}

	for _, feature := range features {
		if err := generateFeatureIndex(feature); err != nil {
			errorf("Warning: failed to generate index for %s: %v\n", feature, err)
		}
	}

	return nil
}

func generateFeatureIndex(feature string) error {
	// Find feature location
	var status string
	var specDir string
	for _, s := range []string{"active", "backlog", "completed", "archived"} {
		path := filepath.Join(specsDir, s, feature)
		if _, err := os.Stat(path); err == nil {
			status = s
			specDir = path
			break
		}
	}

	if status == "" {
		return fmt.Errorf("feature not found")
	}

	// Read MRD for title
	var title string
	mrdPath := filepath.Join(specDir, "mrd.json")
	if data, err := os.ReadFile(mrdPath); err == nil {
		var mrd types.MRD
		if err := json.Unmarshal(data, &mrd); err == nil {
			title = mrd.Metadata.Title
		}
	}
	if title == "" {
		title = feature
	}

	// Create docs directory
	docsFeatureDir := filepath.Join(docsDir, "design", feature)
	if err := os.MkdirAll(docsFeatureDir, 0755); err != nil {
		return err
	}

	// Generate index.md
	var sb strings.Builder
	_, _ = fmt.Fprintf(&sb, "# %s\n\n", title)
	_, _ = fmt.Fprintf(&sb, "**Status:** %s\n\n", status)

	sb.WriteString("## Specs\n\n")

	// Check which specs exist
	specs := []struct {
		file string
		name string
	}{
		{"mrd.json", "Market Requirements Document (MRD)"},
		{"prd.json", "Product Requirements Document (PRD)"},
		{"evaluation.json", "Spec Evaluation"},
	}

	for _, spec := range specs {
		specPath := filepath.Join(specDir, spec.file)
		if _, err := os.Stat(specPath); err == nil {
			docFile := strings.TrimSuffix(spec.file, ".json") + ".md"
			_, _ = fmt.Fprintf(&sb, "- [%s](%s) | [spec](../../specs/%s/%s/%s)\n",
				spec.name, docFile, status, feature, spec.file)
		}
	}

	indexPath := filepath.Join(docsFeatureDir, "index.md")
	if err := os.WriteFile(indexPath, []byte(sb.String()), 0644); err != nil {
		return err
	}

	printf("Generated %s\n", indexPath)
	return nil
}
