# PRD Schema

Product Requirements Document schema.

## Go Type

```go
type PRD struct {
    Metadata                  DocumentMetadata          `json:"metadata"`
    MRDReference              string                    `json:"mrd_reference"`
    Overview                  PRDOverview               `json:"overview"`
    FunctionalRequirements    []FunctionalRequirement   `json:"functional_requirements"`
    NonFunctionalRequirements []NonFunctionalRequirement `json:"non_functional_requirements,omitempty"`
    Checkpoints               []Checkpoint              `json:"checkpoints,omitempty"`
    TechnicalDesign           *TechnicalDesign          `json:"technical_design,omitempty"`
    Dependencies              []Dependency              `json:"dependencies,omitempty"`
    RolloutPlan               *RolloutPlan              `json:"rollout_plan,omitempty"`
}
```

## Fields

### functional_requirements

Detailed, implementable requirements.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier (e.g., `FR-001`) |
| `mrd_requirement` | string | No | Reference to MRD requirement |
| `description` | string | Yes | What needs to be implemented |
| `acceptance_criteria` | string[] | Yes | Testable criteria |
| `test_hints` | TestHint[] | No | Guidance for tests |
| `uncertainty` | enum | No | `low`, `medium`, `high` |
| `uncertainty_reason` | string | No | Why uncertainty exists |
| `discovery_prompt` | string | No | Question to resolve |
| `status` | enum | No | `pending`, `in_progress`, `blocked`, `implemented`, `validated` |
| `blocked_reason` | string | No | Why blocked |
| `checkpoint` | string | No | Checkpoint ID reference |

### acceptance_criteria

Array of strings. Each criterion should be:

- Specific and unambiguous
- Testable (can verify programmatically)
- Complete (covers success, failure, edge cases)

### test_hints

Guidance for test generation.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | enum | Yes | `unit`, `integration`, `edge_case`, `performance`, `security` |
| `focus` | string | Yes | What to test |

### non_functional_requirements

Performance, security, and operational requirements.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier (e.g., `NFR-001`) |
| `category` | enum | Yes | `performance`, `security`, `availability`, `scalability`, `maintainability` |
| `description` | string | Yes | The requirement |
| `measurement` | string | Yes | How to measure |
| `priority` | enum | Yes | `must_have`, `should_have`, `could_have` |

### checkpoints

Validation gates during implementation.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier (e.g., `CP-001`) |
| `name` | string | Yes | Human-readable name |
| `after_requirements` | string[] | Yes | Requirement IDs that trigger |
| `validation_type` | enum | Yes | `human_review`, `automated`, `skip` |
| `validation_criteria` | string[] | No | For automated validation |
| `pause_on_discovery` | boolean | No | Stop if discoveries emerge |
| `description` | string | No | Additional context |

### Uncertainty Levels

| Level | Meaning | Action |
|-------|---------|--------|
| `low` | Well understood | Implement directly |
| `medium` | Some unknowns | Clarify during implementation |
| `high` | Significant unknowns | Resolve before implementation |

When `uncertainty` is `high`, include:

- `uncertainty_reason`: Why it's uncertain
- `discovery_prompt`: Question to answer

## Example

See [example-prd.json](https://github.com/plexusone/dark-factory-project-spec/blob/main/examples/example-prd.json).
