# Dark Factory Project Spec

A specification system for creating complete, unambiguous requirements documents (MRD/PRD) that can be executed autonomously by AI development systems.

## The Problem

AI development systems fail when specs are incomplete or ambiguous. The bottleneck isn't AI capability—it's spec quality.

Common failure modes:

- **Missing edge cases** - AI makes assumptions that differ from intent
- **Ambiguous requirements** - Multiple valid interpretations lead to wrong choices
- **No validation gates** - Errors compound without course correction
- **Unknown unknowns** - Neither human nor AI realizes what's missing

## The Solution

Dark Factory provides:

```mermaid
graph LR
    A[Human Intent] --> B[Structured Spec]
    B --> C[Validation]
    C --> D[AI Execution]
    D --> E[Checkpoints]
    E --> F[Human Review]
    F --> D
    F --> G[Completion]
```

1. **Structured spec formats** - JSON schemas for MRD, PRD, Constitution
2. **Guided spec creation** - Claude skill that interviews humans for completeness
3. **Validation gates** - GO/NO-GO decisions before execution starts
4. **Checkpoints** - Mid-execution pauses for course correction
5. **Uncertainty markers** - Explicit flagging of unknowns

## Quick Example

```json
{
  "metadata": {
    "id": "PRD-2025-001",
    "title": "User Authentication",
    "feature": "user-authentication",
    "status": "approved"
  },
  "functional_requirements": [
    {
      "id": "FR-001",
      "description": "Users can log in with email and password",
      "acceptance_criteria": [
        "Valid credentials return JWT token",
        "Invalid credentials return 401 with error message",
        "Account lockout after 5 failed attempts"
      ],
      "uncertainty": "low",
      "checkpoint": "CP-001"
    }
  ],
  "checkpoints": [
    {
      "id": "CP-001",
      "after_requirements": ["FR-001"],
      "validation": "human_review",
      "pause_on_discovery": true
    }
  ]
}
```

## Key Features

<div class="grid cards" markdown>

-   :material-file-document-outline:{ .lg .middle } __Structured Specs__

    ---

    Go types generate JSON Schema for type-safe, validated specs

    [:octicons-arrow-right-24: Learn more](concepts/overview.md)

-   :material-checkbox-marked-outline:{ .lg .middle } __Checkpoints__

    ---

    Pause execution at defined points for validation and discovery

    [:octicons-arrow-right-24: Learn more](concepts/checkpoints.md)

-   :material-help-circle-outline:{ .lg .middle } __Uncertainty Markers__

    ---

    Explicitly flag unknowns instead of hoping AI figures it out

    [:octicons-arrow-right-24: Learn more](concepts/prd.md#uncertainty)

-   :material-gavel:{ .lg .middle } __Evaluation Gates__

    ---

    GO/NO-GO decisions with structured findings before execution

    [:octicons-arrow-right-24: Learn more](concepts/evaluation.md)

</div>

## Getting Started

```bash
# Install the CLI
go install github.com/plexusone/dark-factory-project-spec/cmd/dfspec@latest

# Initialize a new feature
dfspec init my-feature

# List all features
dfspec list
```

[:octicons-arrow-right-24: Installation Guide](getting-started/installation.md)
