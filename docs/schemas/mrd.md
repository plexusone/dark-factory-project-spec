# MRD Schema

Market Requirements Document schema.

## Go Type

```go
type MRD struct {
    Metadata              DocumentMetadata       `json:"metadata"`
    ConstitutionRef       string                 `json:"constitution_ref,omitempty"`
    ConstitutionOverrides *ConstitutionOverride  `json:"constitution_overrides,omitempty"`
    ProblemStatement      ProblemStatement       `json:"problem_statement"`
    TargetUsers           []TargetUser           `json:"target_users"`
    SuccessMetrics        []SuccessMetric        `json:"success_metrics"`
    MarketContext         MarketContext          `json:"market_context,omitempty"`
    HighLevelRequirements []HighLevelRequirement `json:"high_level_requirements"`
    Constraints           []Constraint           `json:"constraints,omitempty"`
    NonGoals              []string               `json:"non_goals,omitempty"`
    Assumptions           []string               `json:"assumptions,omitempty"`
    OpenQuestions         []OpenQuestion         `json:"open_questions,omitempty"`
}
```

## Fields

### metadata

Document identification and versioning.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier (e.g., `MRD-2025-001`) |
| `title` | string | Yes | Human-readable title |
| `feature` | string | No | Feature name for organization |
| `version` | string | Yes | Semantic version |
| `status` | enum | Yes | `draft`, `review`, `approved`, `superseded` |
| `authors` | string[] | Yes | List of authors |
| `created_at` | string | Yes | RFC3339 timestamp |
| `updated_at` | string | Yes | RFC3339 timestamp |
| `approved_at` | string | No | RFC3339 timestamp |
| `approved_by` | string[] | No | List of approvers |

### problem_statement

Describes the problem being solved.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `summary` | string | Yes | Brief problem description |
| `current_state` | string | Yes | How things work today |
| `desired_state` | string | Yes | How things should work |
| `impact` | string | Yes | Business impact |

### target_users

User personas.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Persona name |
| `description` | string | Yes | Who this user is |
| `needs` | string[] | Yes | What they need |
| `pain_points` | string[] | No | Current frustrations |
| `priority` | enum | Yes | `primary`, `secondary`, `tertiary` |

### success_metrics

How success is measured.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier |
| `name` | string | Yes | Metric name |
| `description` | string | Yes | What is measured |
| `current_value` | string | No | Baseline value |
| `target_value` | string | Yes | Goal value |
| `measurement_method` | string | Yes | How to measure |
| `timeframe` | string | No | When target should be achieved |

### high_level_requirements

Business-level requirements.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier |
| `description` | string | Yes | What is needed |
| `rationale` | string | Yes | Why it's needed |
| `priority` | enum | Yes | `must_have`, `should_have`, `could_have`, `wont_have` |
| `dependencies` | string[] | No | Other requirement IDs |

### constraints

Limitations and boundaries.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier |
| `type` | enum | Yes | `technical`, `business`, `regulatory`, `resource`, `timeline` |
| `description` | string | Yes | The constraint |
| `rationale` | string | No | Why it exists |
| `negotiable` | boolean | Yes | Can it be changed? |

### open_questions

Unresolved questions.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier |
| `question` | string | Yes | The question |
| `context` | string | No | Background |
| `impact` | string | No | What depends on the answer |
| `owner` | string | No | Who should answer |
| `due_date` | string | No | When answer is needed |
| `answer` | string | No | Resolved answer |
| `answered_at` | string | No | When answered |

## Example

See [example-mrd.json](https://github.com/plexusone/dark-factory-project-spec/blob/main/examples/example-mrd.json).
