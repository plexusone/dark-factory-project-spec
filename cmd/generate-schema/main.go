// Package main generates JSON Schema files from Go types.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/plexusone/dark-factory-project-spec/types"
)

// schemaConfig defines a schema to generate.
type schemaConfig struct {
	Type     any
	Filename string
}

func main() {
	schemas := []schemaConfig{
		{Type: types.MRD{}, Filename: "mrd.schema.json"},
		{Type: types.PRD{}, Filename: "prd.schema.json"},
		{Type: types.SpecEvaluation{}, Filename: "evaluation.schema.json"},
		{Type: types.ExecutionReport{}, Filename: "execution-report.schema.json"},
		{Type: types.EvaluationRubric{}, Filename: "evaluation-rubric.schema.json"},
		{Type: types.Constitution{}, Filename: "constitution.schema.json"},
		{Type: types.Roadmap{}, Filename: "roadmap.schema.json"},
		{Type: types.ExecutionState{}, Filename: "execution-state.schema.json"},
	}

	outputDir := "schemas"
	if len(os.Args) > 1 {
		outputDir = os.Args[1]
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	reflector := &jsonschema.Reflector{
		DoNotReference: false,
		ExpandedStruct: false,
	}

	for _, cfg := range schemas {
		schema := reflector.Reflect(cfg.Type)

		data, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling schema for %s: %v\n", cfg.Filename, err)
			os.Exit(1)
		}

		outputPath := filepath.Join(outputDir, cfg.Filename)
		if err := os.WriteFile(outputPath, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing schema to %s: %v\n", outputPath, err)
			os.Exit(1)
		}

		fmt.Printf("Generated %s\n", outputPath)
	}

	fmt.Println("Schema generation complete.")
}
