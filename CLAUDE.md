# CLAUDE.md

Project-specific instructions for dark-factory-project-spec.

## Project Overview

This project implements a "dark factory" spec system - a framework for creating complete, unambiguous requirements documents (MRD/PRD) that can be executed autonomously by AI development systems.

## Key Components

- **types/**: Go structs defining MRD, PRD, checkpoint, and evaluation types (source of truth)
- **schemas/**: Generated JSON schemas from Go types
- **cmd/dfspec/**: CLI tool for managing specs
- **skills/dfspec-guide/**: Claude skill for guiding users through spec creation
- **skills/dfspec-exec/**: Claude skill for executing approved specs with checkpoints
- **examples/**: Working examples of MRD, PRD, and evaluation outputs
- **specs/**: Project specs organized by status (active, backlog, completed, archived)
- **docs/**: MkDocs documentation site

## CLI Commands (dfspec)

### Install CLI

```bash
go install ./cmd/dfspec
# or
make install-cli
```

### List features

```bash
dfspec list              # All features grouped by status
dfspec list active       # Filter by status
dfspec list --json       # JSON output
```

### Initialize new feature

```bash
dfspec init my-feature                    # Creates in backlog/
dfspec init my-feature --status active    # Creates in active/
dfspec init my-feature --author "Name"    # Set author
```

### Move between statuses

```bash
dfspec move feature-name active           # Start working
dfspec move feature-name completed -v v1.0.0  # Mark done with version
dfspec move feature-name archived -r "Reason"  # Archive with reason
```

### Validate specs

```bash
dfspec validate                # Validate all
dfspec validate feature-name   # Validate specific feature
```

### Generate documentation

```bash
dfspec generate roadmap        # Generate ROADMAP.md
dfspec generate index          # Generate index.md for all features
dfspec generate index feature  # Generate for specific feature
```

### Execution management

```bash
dfspec execute status feature  # Check execution status
dfspec execute start feature   # Initialize execution state
dfspec execute log feature     # View execution log
dfspec execute pause feature   # Pause execution
```

### Checkpoint management

```bash
dfspec checkpoint list feature              # List checkpoints
dfspec checkpoint show feature CP-001       # Show checkpoint details
dfspec checkpoint approve feature CP-001    # Approve a checkpoint
dfspec checkpoint reject feature CP-001 -m "reason"  # Reject
```

## Development Commands

```bash
# Build and test
go build ./...
go test ./...

# Generate schemas from Go types
go run cmd/generate-schema/main.go
# or
make generate-schema

# Lint
golangci-lint run

# Full check
make all
```

## MkDocs Documentation

```bash
# Install dependencies
pip install -r requirements.txt

# Serve locally
mkdocs serve

# Build static site
mkdocs build
```

## Go-First Schema Approach

This project follows a **Go-first** approach where Go structs are the source of truth:

1. Define/update types in `types/*.go`
2. Run `go run cmd/generate-schema/main.go` to regenerate schemas
3. Validate examples against schemas

**Never hand-edit** the JSON schema files in `schemas/`.

## Skills

Two skills are provided:

### /dfspec-guide
Creates complete MRD/PRD documents interactively.

### /dfspec-exec
Executes PRD requirements with checkpoint-based human oversight.

Install both skills:
```bash
make install-skill
```

Or individually:
```bash
make install-skill-dfspec-guide
make install-skill-dfspec-exec
```

Skill structure:
```
skills/{skill-name}/
├── SKILL.md      # Main skill definition
├── references/   # Templates, guides
└── agents/       # Specialized agents
```

## Project Structure

```
specs/
├── roadmap.json          # Feature priorities (auto-updated by CLI)
├── constitution.json     # Organization standards
├── active/               # Features being worked on
│   └── {feature}/
│       ├── mrd.json
│       └── prd.json
├── backlog/              # Planned features
├── completed/            # Shipped features
└── archived/             # Cancelled features
```

## Schema Changes

When modifying types:

1. Update the Go struct in `types/*.go`
2. Add appropriate `json` tags
3. Run schema generation: `make generate-schema`
4. Update examples if needed
5. Run tests: `go test ./...`

## Feature Lifecycle

1. **Create**: `dfspec init feature-name` → backlog/
2. **Activate**: `dfspec move feature-name active` → active/
3. **Complete**: `dfspec move feature-name completed -v v1.0.0` → completed/
4. (or) **Archive**: `dfspec move feature-name archived -r "reason"` → archived/

The CLI automatically updates `roadmap.json` on every move.

## Evaluation Integration

The evaluation types in `types/evaluation.go` are compatible with `structured-evaluation`. The output format matches what sevaluation expects.
