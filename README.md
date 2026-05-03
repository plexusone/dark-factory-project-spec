# Dark Factory Project Spec

A specification system for creating complete, unambiguous requirements documents (MRD/PRD) that can be executed autonomously by AI development systems.

## Overview

Dark Factory addresses the gap between human intent and AI execution by providing:

- **Structured spec formats** - Go types that generate JSON Schema for MRD, PRD, Constitution, and Roadmap documents
- **Checkpoint system** - Mid-execution validation points with discovery handling
- **Uncertainty markers** - Explicit flagging of unknowns in requirements
- **Evaluation integration** - GO/NO-GO gating via structured evaluation
- **CLI tooling** - `dfspec` for managing specs through their lifecycle

## Installation

```bash
# Install the CLI
go install github.com/plexusone/dark-factory-project-spec/cmd/dfspec@latest

# Or build from source
git clone https://github.com/plexusone/dark-factory-project-spec.git
cd dark-factory-project-spec
make install-cli
```

## Quick Start

```bash
# Initialize a new feature in backlog
dfspec init user-authentication

# List all features
dfspec list

# Move feature to active when ready to work
dfspec move user-authentication active

# Validate specs
dfspec validate user-authentication

# Generate human-readable roadmap
dfspec generate roadmap
```

## Project Structure

```
specs/
├── roadmap.json          # Feature priorities and status
├── constitution.json     # Organization-wide standards
├── active/               # Features being worked on
│   └── {feature}/
│       ├── mrd.json      # Market Requirements Document
│       ├── prd.json      # Product Requirements Document
│       └── evaluation.json
├── backlog/              # Planned features
├── completed/            # Shipped features
└── archived/             # Cancelled/superseded features

docs/
└── design/
    └── {feature}/        # Human-readable docs per feature
```

## Spec Types

### Constitution

Organization-wide source of truth for:

- Tech stack (languages, frameworks, databases)
- Architecture standards (patterns, principles, API design)
- File organization (specs/, docs/ structure)
- Quality standards (testing, code review)
- Security standards (auth, data protection)
- Process standards (branching, releases)

### MRD (Market Requirements Document)

High-level business requirements:

- Problem statement (current state, desired state, impact)
- Target users (personas, needs, pain points)
- Success metrics (KPIs, measurement methods)
- High-level requirements (must-have, should-have, etc.)
- Constraints and non-goals

### PRD (Product Requirements Document)

Detailed implementation requirements:

- Functional requirements with acceptance criteria
- Non-functional requirements (performance, security)
- Test hints for comprehensive coverage
- Checkpoint references for validation gates
- Uncertainty markers for unknowns

### Roadmap

Feature tracking with status-based organization:

- **Active** - Currently being implemented
- **Backlog** - Planned, prioritized by importance
- **Completed** - Shipped and released
- **Archived** - Cancelled with reason

## CLI Commands

| Command | Description |
|---------|-------------|
| `dfspec list [status]` | List features, optionally filtered by status |
| `dfspec init <feature>` | Create new feature in backlog with MRD template |
| `dfspec move <feature> <status>` | Move feature between statuses |
| `dfspec validate [feature]` | Validate specs against schemas |
| `dfspec generate roadmap` | Generate ROADMAP.md from roadmap.json |
| `dfspec generate index [feature]` | Generate index.md for feature docs |

## Claude Code Skills

Two Claude Code skills are provided:

### /dfspec-guide
Guides users through creating complete MRD/PRD specs interactively.

### /dfspec-exec
Executes approved PRD specs autonomously with checkpoint-based human oversight.

```bash
# Install the skills
make install-skill

# Use in Claude Code
/dfspec-guide    # Create a new spec
/dfspec-exec     # Execute an approved spec
```

The skill provides:

- Interactive interview for gathering requirements
- Probing questions for edge cases and unknowns
- Draft generation in JSON format
- Evaluation against completeness rubric
- Human review gate before development

## Development

```bash
# Build
make build

# Test
make test

# Lint
make lint

# Generate JSON schemas from Go types
make generate-schema

# Validate examples against schemas
make validate-examples
```

## Schema Generation

Go types are the source of truth. JSON schemas are generated:

```bash
go run cmd/generate-schema/main.go
```

Schemas are output to `schemas/`:

- `mrd.schema.json`
- `prd.schema.json`
- `constitution.schema.json`
- `roadmap.schema.json`
- `evaluation.schema.json`

## Key Concepts

### Checkpoints

Specs define validation points where execution pauses:

```json
{
  "checkpoints": [
    {
      "id": "CP-001",
      "after_requirements": ["FR-001", "FR-002"],
      "validation": "human_review",
      "pause_on_discovery": true
    }
  ]
}
```

### Uncertainty Markers

Requirements can flag unknowns explicitly:

```json
{
  "id": "FR-003",
  "description": "Handle concurrent uploads",
  "uncertainty": "high",
  "uncertainty_reason": "Concurrency model not specified",
  "discovery_prompt": "What happens when two users upload simultaneously?"
}
```

### Constitution Overrides

Individual specs can override constitution defaults with justification:

```json
{
  "constitution_ref": "specs/constitution.json",
  "constitution_overrides": {
    "overrides": [
      {
        "path": "tech_stack.languages",
        "original_value": "Go",
        "override_value": "Rust",
        "justification": "Performance-critical component requires Rust"
      }
    ]
  }
}
```

## License

MIT
