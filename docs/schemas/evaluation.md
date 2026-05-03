# Evaluation Schema

Spec evaluation schema.

## Go Type

```go
type SpecEvaluation struct {
    Metadata   EvaluationMetadata  `json:"metadata"`
    Subject    EvaluationSubject   `json:"subject"`
    Categories []CategoryScore     `json:"categories"`
    Findings   []Finding           `json:"findings"`
    Summary    EvaluationSummary   `json:"summary"`
    Decision   Decision            `json:"decision"`
}
```

## Fields

### metadata

Evaluation identification.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `evaluation_id` | string | Yes | Unique identifier |
| `evaluated_at` | string | Yes | RFC3339 timestamp |
| `evaluator` | string | Yes | Who/what evaluated |
| `rubric_version` | string | Yes | Rubric version used |
| `model_id` | string | No | Model used |
| `judge_runs` | JudgeRun[] | No | Multi-judge runs |

### subject

What was evaluated.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | enum | Yes | `mrd`, `prd` |
| `id` | string | Yes | Document ID |
| `version` | string | Yes | Document version |
| `title` | string | Yes | Document title |

### categories

Scores by category.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `category` | string | Yes | Category name |
| `weight` | number | Yes | Weight (0-1) |
| `score` | number | Yes | Score (0-10) |
| `rationale` | string | Yes | Explanation |

Standard categories:

| Category | Weight | Description |
|----------|--------|-------------|
| `problem_definition` | 0.15 | Clear problem statement |
| `requirements_clarity` | 0.25 | Unambiguous requirements |
| `acceptance_criteria` | 0.25 | Testable criteria |
| `constraints` | 0.10 | Documented limitations |
| `non_goals` | 0.10 | Explicit scope boundaries |
| `uncertainty_handling` | 0.15 | Flagged unknowns |

### findings

Specific issues found.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier |
| `severity` | enum | Yes | `critical`, `high`, `medium`, `low`, `info` |
| `category` | string | Yes | Related category |
| `title` | string | Yes | Short description |
| `description` | string | Yes | Full details |
| `location` | string | No | Where in spec |
| `suggestion` | string | No | How to fix |
| `blocking` | boolean | Yes | Blocks approval? |

### summary

Overall assessment.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `overall_score` | number | Yes | Weighted average (0-10) |
| `strengths` | string[] | Yes | What's done well |
| `weaknesses` | string[] | Yes | Areas for improvement |
| `finding_counts` | FindingCounts | Yes | Count by severity |

### decision

GO/NO-GO recommendation.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | enum | Yes | `go`, `no_go`, `conditional` |
| `rationale` | string | Yes | Explanation |
| `conditions` | string[] | No | Conditions for approval |
| `requires_human_review` | boolean | Yes | Needs human review? |
| `human_review_reason` | string | No | Why human review needed |

## Example

```json
{
  "metadata": {
    "evaluation_id": "EVAL-2025-001",
    "evaluated_at": "2025-01-15T10:00:00Z",
    "evaluator": "spec-guide-reviewer",
    "rubric_version": "1.0.0"
  },
  "subject": {
    "type": "prd",
    "id": "PRD-2025-001",
    "version": "1.0.0",
    "title": "User Authentication"
  },
  "categories": [
    {
      "category": "requirements_clarity",
      "weight": 0.25,
      "score": 8.0,
      "rationale": "Clear requirements with minor edge case gaps"
    }
  ],
  "findings": [
    {
      "id": "F-001",
      "severity": "medium",
      "category": "acceptance_criteria",
      "title": "Missing error format",
      "description": "FR-001 doesn't specify error response format",
      "suggestion": "Add error response schema",
      "blocking": false
    }
  ],
  "summary": {
    "overall_score": 7.8,
    "strengths": ["Clear problem definition"],
    "weaknesses": ["Some edge cases missing"],
    "finding_counts": {
      "critical": 0,
      "high": 0,
      "medium": 1,
      "low": 0,
      "info": 0
    }
  },
  "decision": {
    "type": "go",
    "rationale": "Spec meets quality bar",
    "requires_human_review": false
  }
}
```

See [example-evaluation.json](https://github.com/plexusone/dark-factory-project-spec/blob/main/examples/example-evaluation.json).
