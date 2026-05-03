# Installation

## Prerequisites

- Go 1.21 or later
- Git

## Install the CLI

### From Source (Recommended)

```bash
git clone https://github.com/plexusone/dark-factory-project-spec.git
cd dark-factory-project-spec
make install-cli
```

### Using Go Install

```bash
go install github.com/plexusone/dark-factory-project-spec/cmd/dfspec@latest
```

### Verify Installation

```bash
dfspec --version
# dfspec version 0.1.0
```

## Install the Claude Skills

The `/dfspec-guide` skill helps create complete specs interactively:

```bash
# From the repo directory
make install-skill
```

This creates symlinks at `~/.claude/skills/dfspec-guide` and `~/.claude/skills/dfspec-exec`.

## Project Setup

Initialize a new project with the spec structure:

```bash
mkdir my-project
cd my-project

# Create directory structure
mkdir -p specs/{active,backlog,completed,archived}

# Create initial constitution
cat > specs/constitution.json << 'EOF'
{
  "metadata": {
    "version": "1.0.0",
    "effective_date": "2025-01-01T00:00:00Z",
    "approved_by": ["Your Name"]
  },
  "tech_stack": {
    "languages": [
      {"name": "Go", "status": "preferred"}
    ]
  }
}
EOF

# Create initial roadmap
cat > specs/roadmap.json << 'EOF'
{
  "version": "1.0.0",
  "updated_at": "2025-01-01T00:00:00Z",
  "active": [],
  "backlog": [],
  "completed": [],
  "archived": []
}
EOF
```

## Development Setup

For contributing to the project:

```bash
git clone https://github.com/plexusone/dark-factory-project-spec.git
cd dark-factory-project-spec

# Install dependencies
go mod download

# Build
make build

# Run tests
make test

# Run linter
make lint
```

## Optional: Schema Validation

To validate specs against JSON schemas, install `ajv-cli`:

```bash
npm install -g ajv-cli

# Validate examples
make validate-examples
```
