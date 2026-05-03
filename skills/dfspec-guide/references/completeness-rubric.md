# Completeness Rubric

This rubric defines how to evaluate MRD and PRD documents for completeness. It follows the structured-evaluation format for compatibility.

## Overview

Each category is scored 0-10:
- **10**: Exemplary - exceeds expectations, no improvements needed
- **8-9**: Strong - meets expectations, minor improvements possible
- **6-7**: Adequate - acceptable with some gaps
- **4-5**: Weak - significant gaps need addressing
- **2-3**: Poor - major issues, requires substantial revision
- **0-1**: Missing - section absent or unusable

## MRD Evaluation Rubric

### Categories and Weights

| Category | Weight | Description |
|----------|--------|-------------|
| Problem Definition | 20% | Clarity and specificity of the problem |
| Target Users | 15% | User personas and needs definition |
| Success Metrics | 15% | Measurable outcomes |
| High-Level Requirements | 20% | Requirement coverage and prioritization |
| Constraints & Assumptions | 15% | Documented limitations |
| Non-Goals | 15% | Explicit scope boundaries |

### Problem Definition (20%)

**Full marks (10):**
- Problem clearly stated in specific terms
- Current state described with concrete examples
- Desired state clearly articulated
- Business impact quantified or well-explained
- Not conflated with a solution

**Deduct points for:**
- Vague problem description (-2)
- No examples or specifics (-2)
- Solution masquerading as problem (-3)
- No impact statement (-2)
- Missing current/desired state (-2)

**Score guide:**
- 10: "Search queries on datasets >1M rows take 30+ seconds, blocking analysts from iterating on reports"
- 5: "Search is slow"
- 2: "We need a caching layer"

### Target Users (15%)

**Full marks (10):**
- 2-4 distinct personas identified
- Clear priority (primary/secondary/tertiary)
- Specific needs per persona
- Pain points documented
- No "everyone" or "all users"

**Deduct points for:**
- No personas defined (-5)
- Only one persona when multiple exist (-2)
- Features listed as needs (-2)
- "Everyone" as target user (-3)
- Missing priority designation (-2)

### Success Metrics (15%)

**Full marks (10):**
- 3-5 measurable metrics
- Specific targets (not "improve X")
- Baseline values included
- Measurement method described
- Mix of leading and lagging indicators

**Deduct points for:**
- No metrics (-5)
- Unmeasurable metrics (-3)
- No targets (-2)
- No measurement method (-2)
- Vanity metrics only (-2)

### High-Level Requirements (20%)

**Full marks (10):**
- All business capabilities covered
- MoSCoW prioritization applied
- Dependencies identified
- Rationale for each requirement
- Clear scope (not overly broad)

**Deduct points for:**
- Missing capabilities (-3)
- No prioritization (-3)
- Overly broad requirements (-2)
- No rationale (-2)
- Missing dependencies (-2)

### Constraints & Assumptions (15%)

**Full marks (10):**
- Technical constraints documented
- Business constraints documented
- Regulatory/compliance noted
- Assumptions explicit
- Negotiability indicated

**Deduct points for:**
- No constraints (-3)
- No assumptions (-2)
- Missing obvious constraints (-2)
- Unlabeled negotiability (-2)

### Non-Goals (15%)

**Full marks (10):**
- 3-5 explicit non-goals
- Addresses likely assumptions
- Distinguishes "not now" vs "never"
- Prevents obvious scope creep

**Deduct points for:**
- No non-goals (-5)
- Vague non-goals (-2)
- Missing obvious exclusions (-2)

---

## PRD Evaluation Rubric

### Categories and Weights

| Category | Weight | Description |
|----------|--------|-------------|
| Requirements Clarity | 25% | Clear, unambiguous requirements |
| Acceptance Criteria | 25% | Testable Given-When-Then criteria |
| Test Coverage | 15% | Test hints and edge cases |
| Dependencies | 15% | External dependency documentation |
| Uncertainty Handling | 10% | Flagged unknowns and discovery prompts |
| Non-Goals | 10% | Explicit scope boundaries |

### Requirements Clarity (25%)

**Full marks (10):**
- Each requirement is specific and testable
- Requirements describe "what" not "how"
- No ambiguous language ("should", "may", "might")
- Traceability to MRD requirements
- Clear requirement IDs

**Deduct points for:**
- Ambiguous requirements (-2 each, max -6)
- Implementation details instead of requirements (-2)
- Missing MRD traceability (-2)
- No requirement IDs (-2)

**Severity of findings:**
- Ambiguous requirement: **High** (blocking)
- Missing traceability: **Medium** (human review)
- Style issues: **Low** (auto-approve)

### Acceptance Criteria (25%)

**Full marks (10):**
- Given-When-Then format used
- Happy path covered
- Error paths covered
- Edge cases covered
- Permission scenarios covered

**Deduct points for:**
- Missing acceptance criteria (-5)
- No error cases (-2)
- No edge cases (-2)
- Vague criteria (-2)
- Not testable as written (-2)

**Severity of findings:**
- Missing acceptance criteria: **Critical** (blocking)
- Missing error cases: **High** (blocking)
- Missing edge cases: **Medium** (human review)

### Test Coverage (15%)

**Full marks (10):**
- Unit test hints provided
- Integration test hints provided
- Edge cases identified
- Security considerations noted
- Performance considerations noted

**Deduct points for:**
- No test hints (-3)
- Missing unit test guidance (-2)
- No edge case hints (-2)
- No security considerations (-2)

### Dependencies (15%)

**Full marks (10):**
- All external dependencies listed
- Dependency status known
- Owners identified
- Versions specified where applicable
- Risk of unavailability noted

**Deduct points for:**
- Missing dependencies (-3)
- Unknown dependency status (-2)
- No owners (-2)
- Missing version requirements (-2)

### Uncertainty Handling (10%)

**Full marks (10):**
- Uncertainty levels assigned to all requirements
- High-uncertainty items have reasons
- Discovery prompts provided for unknowns
- Checkpoints aligned with uncertainty

**Deduct points for:**
- No uncertainty markers (-3)
- High uncertainty without reason (-2)
- No discovery prompts (-2)
- Checkpoints ignore uncertainty (-2)

### Non-Goals (10%)

**Full marks (10):**
- Technical non-goals stated
- Features explicitly excluded
- Clear boundaries for V1

**Deduct points for:**
- No non-goals (-3)
- Ambiguous scope (-2)

---

## Finding Severity Mappings

### Critical (blocks approval)
- Missing acceptance criteria
- Requirements cannot be tested
- Conflicting requirements
- Missing critical integration details

### High (blocks approval)
- Ambiguous requirements
- Missing error handling requirements
- No success metrics
- Missing dependencies

### Medium (requires human review)
- Incomplete edge cases
- Missing test hints
- Low uncertainty markers on unclear requirements
- Minor ambiguities

### Low (auto-approve)
- Style inconsistencies
- Missing optional fields
- Minor documentation gaps

### Info (auto-approve)
- Suggestions for improvement
- Best practice recommendations
- Optional enhancements

---

## Pass Criteria

### MRD Pass Criteria
```json
{
  "min_score": 7.5,
  "max_critical": 0,
  "max_high": 0,
  "max_medium": -1,
  "human_review_required": ["medium_findings_exist"]
}
```

### PRD Pass Criteria
```json
{
  "min_score": 7.5,
  "max_critical": 0,
  "max_high": 0,
  "max_medium": -1,
  "human_review_required": ["medium_findings_exist", "high_uncertainty_requirements"]
}
```

### Decision Logic

1. **Critical or High findings exist** → NO-GO (must fix)
2. **Score < 7.0** → NO-GO (must improve)
3. **Score >= 7.5 AND only Low/Info findings** → GO (auto-approve)
4. **Score 7.0-7.5 OR Medium findings exist** → CONDITIONAL (human review required)

---

## Example Evaluation Output

```json
{
  "categories": [
    {
      "category": "problem_definition",
      "weight": 0.20,
      "score": 9.0,
      "rationale": "Clear problem statement with specific examples and quantified impact"
    },
    {
      "category": "acceptance_criteria",
      "weight": 0.25,
      "score": 6.0,
      "rationale": "Happy paths covered but missing error cases for 3 requirements"
    }
  ],
  "findings": [
    {
      "id": "F-001",
      "severity": "high",
      "category": "acceptance_criteria",
      "title": "Missing error handling criteria",
      "description": "FR-003, FR-005, FR-007 have no acceptance criteria for error scenarios",
      "location": "functional_requirements",
      "suggestion": "Add Given-When-Then criteria for invalid input, system failures, and permission errors",
      "blocking": true
    }
  ],
  "summary": {
    "overall_score": 7.2,
    "strengths": ["Clear problem definition", "Good user personas"],
    "weaknesses": ["Incomplete error handling", "Missing edge cases"],
    "finding_counts": {"critical": 0, "high": 1, "medium": 2, "low": 3, "info": 1}
  },
  "decision": {
    "type": "no_go",
    "rationale": "High severity finding must be resolved before approval",
    "requires_human_review": false
  }
}
```
