package types_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/plexusone/dark-factory-project-spec/types"
)

func TestMRDExample(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "examples", "example-mrd.json"))
	if err != nil {
		t.Fatalf("failed to read example-mrd.json: %v", err)
	}

	var mrd types.MRD
	if err := json.Unmarshal(data, &mrd); err != nil {
		t.Fatalf("failed to parse MRD: %v", err)
	}

	if mrd.Metadata.Title == "" {
		t.Error("MRD title should not be empty")
	}
	if mrd.Metadata.Feature == "" {
		t.Error("MRD feature should not be empty")
	}
	if len(mrd.HighLevelRequirements) == 0 {
		t.Error("MRD should have high-level requirements")
	}
}

func TestPRDExample(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "examples", "example-prd.json"))
	if err != nil {
		t.Fatalf("failed to read example-prd.json: %v", err)
	}

	var prd types.PRD
	if err := json.Unmarshal(data, &prd); err != nil {
		t.Fatalf("failed to parse PRD: %v", err)
	}

	if prd.Metadata.Title == "" {
		t.Error("PRD title should not be empty")
	}
	if len(prd.FunctionalRequirements) == 0 {
		t.Error("PRD should have functional requirements")
	}
}

func TestConstitutionExample(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "examples", "example-constitution.json"))
	if err != nil {
		t.Fatalf("failed to read example-constitution.json: %v", err)
	}

	var constitution types.Constitution
	if err := json.Unmarshal(data, &constitution); err != nil {
		t.Fatalf("failed to parse Constitution: %v", err)
	}

	if constitution.Metadata.Version == "" {
		t.Error("Constitution version should not be empty")
	}
	if len(constitution.TechStack.Languages) == 0 {
		t.Error("Constitution should specify languages")
	}
}

func TestRoadmapExample(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "examples", "example-roadmap.json"))
	if err != nil {
		t.Fatalf("failed to read example-roadmap.json: %v", err)
	}

	var roadmap types.Roadmap
	if err := json.Unmarshal(data, &roadmap); err != nil {
		t.Fatalf("failed to parse Roadmap: %v", err)
	}

	if roadmap.Version == "" {
		t.Error("Roadmap version should not be empty")
	}
	if len(roadmap.Active) == 0 {
		t.Error("Roadmap should have active features")
	}
}

func TestEvaluationExample(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "examples", "example-evaluation.json"))
	if err != nil {
		t.Fatalf("failed to read example-evaluation.json: %v", err)
	}

	var evaluation types.SpecEvaluation
	if err := json.Unmarshal(data, &evaluation); err != nil {
		t.Fatalf("failed to parse Evaluation: %v", err)
	}

	if evaluation.Metadata.EvaluationID == "" {
		t.Error("Evaluation ID should not be empty")
	}
	if evaluation.Decision.Type == "" {
		t.Error("Evaluation decision type should not be empty")
	}
}

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		unmarshal func([]byte) (any, error)
	}{
		{
			name: "MRD",
			file: "example-mrd.json",
			unmarshal: func(data []byte) (any, error) {
				var v types.MRD
				err := json.Unmarshal(data, &v)
				return v, err
			},
		},
		{
			name: "PRD",
			file: "example-prd.json",
			unmarshal: func(data []byte) (any, error) {
				var v types.PRD
				err := json.Unmarshal(data, &v)
				return v, err
			},
		},
		{
			name: "Constitution",
			file: "example-constitution.json",
			unmarshal: func(data []byte) (any, error) {
				var v types.Constitution
				err := json.Unmarshal(data, &v)
				return v, err
			},
		},
		{
			name: "Roadmap",
			file: "example-roadmap.json",
			unmarshal: func(data []byte) (any, error) {
				var v types.Roadmap
				err := json.Unmarshal(data, &v)
				return v, err
			},
		},
		{
			name: "Evaluation",
			file: "example-evaluation.json",
			unmarshal: func(data []byte) (any, error) {
				var v types.SpecEvaluation
				err := json.Unmarshal(data, &v)
				return v, err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("..", "examples", tt.file))
			if err != nil {
				t.Fatalf("failed to read %s: %v", tt.file, err)
			}

			obj, err := tt.unmarshal(data)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			// Re-marshal and unmarshal to verify round-trip
			remarshaled, err := json.Marshal(obj)
			if err != nil {
				t.Fatalf("failed to re-marshal: %v", err)
			}

			_, err = tt.unmarshal(remarshaled)
			if err != nil {
				t.Fatalf("failed to unmarshal after round-trip: %v", err)
			}
		})
	}
}
