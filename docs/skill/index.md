# Claude Skills

Dark Factory provides two Claude Code skills for the spec-driven development workflow.

## Available Skills

### /dfspec-guide

**Purpose:** Create complete MRD/PRD documents interactively.

**Use when:**
- Starting a new feature
- Need to create requirements documents
- Want guided spec creation with validation

**Workflow:**
1. Discovery - Determine spec type and context
2. Interview - Gather requirements through questions
3. Draft - Generate JSON spec
4. Evaluation - Check completeness
5. Review - Get human approval

[:octicons-arrow-right-24: dfspec-guide documentation](dfspec-guide.md)

### /dfspec-exec

**Purpose:** Execute approved PRD specifications autonomously.

**Use when:**
- PRD is approved and ready for implementation
- Want autonomous execution with human oversight
- Need checkpoint-based validation during development

**Workflow:**
1. Load PRD and initialize state
2. Implement requirements in dependency order
3. Pause at checkpoints for review
4. Handle discoveries as they arise
5. Complete when all requirements implemented

[:octicons-arrow-right-24: dfspec-exec documentation](dfspec-exec.md)

## Installation

```bash
# Install both skills
make install-skill

# Or individually
make install-skill-dfspec-guide
make install-skill-dfspec-exec
```

## Full Workflow

```mermaid
graph LR
    A[Idea] --> B[/dfspec-guide]
    B --> C[MRD]
    C --> D[Evaluate]
    D --> E{Pass?}
    E -->|No| B
    E -->|Yes| F[/dfspec-guide]
    F --> G[PRD]
    G --> H[Evaluate]
    H --> I{Pass?}
    I -->|No| F
    I -->|Yes| J[/dfspec-exec]
    J --> K[Implement]
    K --> L[Checkpoint]
    L --> M{Approve?}
    M -->|No| N[Amend]
    N --> J
    M -->|Yes| K
    K --> O[Done]
```

## CLI Support

Both skills are supported by CLI commands:

| Skill | CLI Commands |
|-------|--------------|
| dfspec-guide | `dfspec init`, `dfspec validate` |
| dfspec-exec | `dfspec execute`, `dfspec checkpoint` |

## Skill Files

Skills follow a standard structure:

```
skills/{skill-name}/
├── SKILL.md           # Main skill definition
├── agents/            # Specialized agent prompts
│   ├── agent1.md
│   └── agent2.md
└── references/        # Templates and guides
    ├── template.md
    └── guide.md
```
