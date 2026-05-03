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

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate specs against schemas",
	Long: `Validate all specs in the specs directory.

Checks:
  - JSON syntax validity
  - Schema conformance
  - Cross-references (PRD→MRD, roadmap→features)
  - Constitution compliance`,
	RunE: runValidate,
}

var (
	validateVerbose bool
)

func init() {
	validateCmd.Flags().BoolVarP(&validateVerbose, "verbose", "v", false, "Verbose output")
}

func runValidate(cmd *cobra.Command, args []string) error {
	var errors []string
	var warnings []string

	// Validate roadmap
	roadmapPath := filepath.Join(specsDir, "roadmap.json")
	roadmap, err := validateRoadmap(roadmapPath)
	if err != nil {
		errors = append(errors, fmt.Sprintf("roadmap.json: %v", err))
	} else if validateVerbose {
		printf("✓ roadmap.json\n")
	}

	// Validate constitution
	constitutionPath := filepath.Join(specsDir, "constitution.json")
	if _, err := validateConstitution(constitutionPath); err != nil {
		if os.IsNotExist(err) {
			warnings = append(warnings, "constitution.json not found (optional)")
		} else {
			errors = append(errors, fmt.Sprintf("constitution.json: %v", err))
		}
	} else if validateVerbose {
		printf("✓ constitution.json\n")
	}

	// Validate features in each status directory
	statuses := []string{"active", "backlog", "completed", "archived"}
	featuresSeen := make(map[string]string) // feature -> status

	for _, status := range statuses {
		statusDir := filepath.Join(specsDir, status)
		if _, err := os.Stat(statusDir); os.IsNotExist(err) {
			continue
		}

		entries, err := os.ReadDir(statusDir)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: %v", status, err))
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			feature := entry.Name()
			featureDir := filepath.Join(statusDir, feature)

			// Check for duplicate features across statuses
			if existingStatus, exists := featuresSeen[feature]; exists {
				errors = append(errors, fmt.Sprintf("feature '%s' exists in both %s and %s", feature, existingStatus, status))
			}
			featuresSeen[feature] = status

			// Validate feature name format (kebab-case)
			if !isKebabCase(feature) {
				errors = append(errors, fmt.Sprintf("%s/%s: feature name must be kebab-case", status, feature))
			}

			// Validate MRD if exists
			mrdPath := filepath.Join(featureDir, "mrd.json")
			if _, err := os.Stat(mrdPath); err == nil {
				if err := validateMRD(mrdPath, feature); err != nil {
					errors = append(errors, fmt.Sprintf("%s/%s/mrd.json: %v", status, feature, err))
				} else if validateVerbose {
					printf("✓ %s/%s/mrd.json\n", status, feature)
				}
			}

			// Validate PRD if exists
			prdPath := filepath.Join(featureDir, "prd.json")
			if _, err := os.Stat(prdPath); err == nil {
				if err := validatePRD(prdPath, feature, featureDir); err != nil {
					errors = append(errors, fmt.Sprintf("%s/%s/prd.json: %v", status, feature, err))
				} else if validateVerbose {
					printf("✓ %s/%s/prd.json\n", status, feature)
				}
			}

			// Validate evaluation if exists
			evalPath := filepath.Join(featureDir, "evaluation.json")
			if _, err := os.Stat(evalPath); err == nil {
				if err := validateEvaluation(evalPath); err != nil {
					errors = append(errors, fmt.Sprintf("%s/%s/evaluation.json: %v", status, feature, err))
				} else if validateVerbose {
					printf("✓ %s/%s/evaluation.json\n", status, feature)
				}
			}
		}
	}

	// Cross-reference: roadmap features should exist
	if roadmap != nil {
		for _, f := range roadmap.Active {
			if status, exists := featuresSeen[f.Feature]; !exists {
				errors = append(errors, fmt.Sprintf("roadmap: active feature '%s' not found in specs/active/", f.Feature))
			} else if status != "active" {
				errors = append(errors, fmt.Sprintf("roadmap: feature '%s' is in roadmap.active but found in specs/%s/", f.Feature, status))
			}
		}
		for _, f := range roadmap.Backlog {
			if status, exists := featuresSeen[f.Feature]; !exists {
				errors = append(errors, fmt.Sprintf("roadmap: backlog feature '%s' not found in specs/backlog/", f.Feature))
			} else if status != "backlog" {
				errors = append(errors, fmt.Sprintf("roadmap: feature '%s' is in roadmap.backlog but found in specs/%s/", f.Feature, status))
			}
		}
		for _, f := range roadmap.Completed {
			if status, exists := featuresSeen[f.Feature]; !exists {
				errors = append(errors, fmt.Sprintf("roadmap: completed feature '%s' not found in specs/completed/", f.Feature))
			} else if status != "completed" {
				errors = append(errors, fmt.Sprintf("roadmap: feature '%s' is in roadmap.completed but found in specs/%s/", f.Feature, status))
			}
		}
		for _, f := range roadmap.Archived {
			if status, exists := featuresSeen[f.Feature]; !exists {
				errors = append(errors, fmt.Sprintf("roadmap: archived feature '%s' not found in specs/archived/", f.Feature))
			} else if status != "archived" {
				errors = append(errors, fmt.Sprintf("roadmap: feature '%s' is in roadmap.archived but found in specs/%s/", f.Feature, status))
			}
		}
	}

	// Print results
	if len(warnings) > 0 {
		printf("\nWarnings:\n")
		for _, w := range warnings {
			printf("  ⚠ %s\n", w)
		}
	}

	if len(errors) > 0 {
		printf("\nErrors:\n")
		for _, e := range errors {
			printf("  ✗ %s\n", e)
		}
		return fmt.Errorf("validation failed with %d error(s)", len(errors))
	}

	printf("\n✓ All specs valid\n")
	return nil
}

func validateRoadmap(path string) (*types.Roadmap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var roadmap types.Roadmap
	if err := json.Unmarshal(data, &roadmap); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	// Validate required fields
	if roadmap.Version == "" {
		return nil, fmt.Errorf("missing version")
	}

	return &roadmap, nil
}

func validateConstitution(path string) (*types.Constitution, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var constitution types.Constitution
	if err := json.Unmarshal(data, &constitution); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	return &constitution, nil
}

func validateMRD(path string, expectedFeature string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var mrd types.MRD
	if err := json.Unmarshal(data, &mrd); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	// Validate required fields
	if mrd.Metadata.ID == "" {
		return fmt.Errorf("missing metadata.id")
	}
	if mrd.Metadata.Title == "" {
		return fmt.Errorf("missing metadata.title")
	}

	// Validate feature matches directory
	if mrd.Metadata.Feature != "" && mrd.Metadata.Feature != expectedFeature {
		return fmt.Errorf("metadata.feature '%s' does not match directory '%s'", mrd.Metadata.Feature, expectedFeature)
	}

	return nil
}

func validatePRD(path string, expectedFeature string, featureDir string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var prd types.PRD
	if err := json.Unmarshal(data, &prd); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	// Validate required fields
	if prd.Metadata.ID == "" {
		return fmt.Errorf("missing metadata.id")
	}
	if prd.Metadata.Title == "" {
		return fmt.Errorf("missing metadata.title")
	}

	// Validate feature matches directory
	if prd.Metadata.Feature != "" && prd.Metadata.Feature != expectedFeature {
		return fmt.Errorf("metadata.feature '%s' does not match directory '%s'", prd.Metadata.Feature, expectedFeature)
	}

	// Validate MRD reference exists
	if prd.MRDReference != "" {
		mrdPath := filepath.Join(featureDir, "mrd.json")
		if _, err := os.Stat(mrdPath); os.IsNotExist(err) {
			return fmt.Errorf("references MRD but mrd.json not found in feature directory")
		}
	}

	return nil
}

func validateEvaluation(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var eval types.SpecEvaluation
	if err := json.Unmarshal(data, &eval); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	return nil
}

func isKebabCase(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= '0' && c <= '9' && i > 0 {
			continue
		}
		if c == '-' && i > 0 && i < len(s)-1 {
			continue
		}
		return false
	}
	return !strings.Contains(s, "--")
}
