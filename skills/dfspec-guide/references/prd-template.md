# PRD Template

This template provides the structure for a Product Requirements Document (PRD). A PRD decomposes high-level MRD requirements into detailed functional requirements with acceptance criteria.

## 1. Metadata

```json
{
  "metadata": {
    "id": "PRD-YYYY-NNN",
    "title": "Human-readable title",
    "version": "1.0.0",
    "status": "draft",
    "authors": ["Name"],
    "created_at": "RFC3339 timestamp",
    "updated_at": "RFC3339 timestamp"
  }
}
```

**Guidance:**
- ID format: `PRD-YYYY-NNN` (year + sequential number)
- Link to parent MRD in `mrd_reference`
- Version follows semantic versioning

## 2. MRD Reference

```json
{
  "mrd_reference": "MRD-2025-001"
}
```

Links this PRD to its parent MRD for traceability.

## 3. Overview

**Purpose:** Summarize what this PRD covers and its goals.

```json
{
  "overview": {
    "summary": "What is being built",
    "goals": ["Goal 1 from MRD", "Goal 2"],
    "target_release": "v1.0.0"
  }
}
```

## 4. Functional Requirements

**Purpose:** Detailed, testable requirements with acceptance criteria.

```json
{
  "functional_requirements": [
    {
      "id": "FR-001",
      "mrd_requirement": "HLR-001",
      "title": "Short name",
      "description": "What the system must do",
      "acceptance_criteria": [
        {
          "id": "AC-001",
          "given": "Precondition",
          "when": "Action",
          "then": "Expected outcome"
        }
      ],
      "test_hints": [
        {
          "type": "unit|integration|e2e|edge_case|performance|security",
          "focus": "What to test"
        }
      ],
      "priority": "must_have|should_have|could_have|wont_have",
      "status": "pending",
      "dependencies": ["FR-002"],
      "checkpoint": "CP-001",
      "uncertainty": "low|medium|high",
      "uncertainty_reason": "Why uncertain (if high)",
      "discovery_prompt": "Question to clarify (if uncertain)"
    }
  ]
}
```

### Writing Good Requirements

**ID Format:** `FR-NNN` (functional requirement) or `NFR-NNN` (non-functional)

**Title:** Short, descriptive name (3-8 words)
- Good: "Upload CSV files"
- Bad: "The file upload functionality requirements"

**Description:** What the system must do, not how
- Good: "Users can upload CSV files up to 100MB with automatic encoding detection"
- Bad: "Use the S3 SDK to upload files to the ingest bucket"

### Writing Acceptance Criteria (Gherkin Format)

Use Given-When-Then format for testable criteria:

```
Given: Precondition or initial state
When: Action performed by user or system
Then: Expected outcome or behavior
```

**Examples:**

```json
{
  "given": "A logged-in user with a valid CSV file",
  "when": "The user uploads the file via the upload form",
  "then": "The file is stored and the user sees a success confirmation"
}
```

```json
{
  "given": "A user attempts to upload a 150MB file",
  "when": "The upload is submitted",
  "then": "The system rejects the upload with error 'File exceeds 100MB limit'"
}
```

**Coverage checklist:**
- [ ] Happy path (normal success case)
- [ ] Error cases (invalid input, failures)
- [ ] Edge cases (empty, max values, boundaries)
- [ ] Permission cases (unauthorized access)

### Test Hints

Guide automated test generation:

```json
{
  "test_hints": [
    {"type": "unit", "focus": "CSV parser handles malformed headers"},
    {"type": "integration", "focus": "Upload endpoint validates file size before processing"},
    {"type": "edge_case", "focus": "Empty file (0 bytes), file at exactly 100MB limit"},
    {"type": "security", "focus": "Reject files with executable content"}
  ]
}
```

### Uncertainty Handling

When requirements are unclear:

```json
{
  "uncertainty": "high",
  "uncertainty_reason": "Concurrent upload behavior not specified",
  "discovery_prompt": "What happens when two users upload the same filename simultaneously?"
}
```

**Uncertainty levels:**
- `low`: Confident in the requirement
- `medium`: Some details unclear, but can proceed
- `high`: Significant unknowns, may need checkpoint validation

## 5. Non-Functional Requirements

**Purpose:** Quality attributes (performance, security, etc.)

```json
{
  "non_functional_requirements": [
    {
      "id": "NFR-001",
      "category": "performance|scalability|reliability|security|usability|maintainability|accessibility",
      "description": "The requirement",
      "measure": "How to measure",
      "target": "Specific target",
      "priority": "must_have"
    }
  ]
}
```

**Examples:**

```json
{
  "id": "NFR-001",
  "category": "performance",
  "description": "API response time under normal load",
  "measure": "95th percentile latency",
  "target": "< 500ms",
  "priority": "must_have"
}
```

## 6. Checkpoints

**Purpose:** Define validation points during execution where the factory pauses.

```json
{
  "checkpoints": [
    {
      "id": "CP-001",
      "name": "Core Upload Complete",
      "description": "Validate basic upload functionality before building dependent features",
      "after_requirements": ["FR-001", "FR-002"],
      "validation": "human_review|automated|skip",
      "pause_on_discovery": true,
      "validation_criteria": ["All tests pass", "Performance meets NFR-001"],
      "status": "pending"
    }
  ]
}
```

**Validation types:**
- `human_review`: Requires human approval to continue
- `automated`: Runs automated validation (tests, linting)
- `skip`: No validation (use for low-risk groupings)

**When to add checkpoints:**
- After critical foundation requirements
- Before building dependent features
- At integration boundaries
- Before irreversible changes (migrations, API versioning)

## 7. Technical Constraints

**Purpose:** Implementation constraints from architecture, existing systems, or decisions.

```json
{
  "technical_constraints": [
    {
      "id": "TC-001",
      "area": "Database",
      "constraint": "Must use existing PostgreSQL cluster",
      "rationale": "Operational team only supports PostgreSQL"
    }
  ]
}
```

## 8. Dependencies

**Purpose:** Track external dependencies that affect implementation.

```json
{
  "dependencies": [
    {
      "id": "DEP-001",
      "name": "Auth Service v2",
      "type": "service|library|api|team|data",
      "description": "Required for user authentication",
      "version": "2.0.0",
      "owner": "Platform Team",
      "status": "available|in_progress|blocked|unknown"
    }
  ]
}
```

## 9. Out of Scope

**Purpose:** Explicitly exclude features to prevent scope creep.

```json
{
  "out_of_scope": [
    "Real-time collaborative editing",
    "Mobile app support",
    "Offline mode"
  ]
}
```

## 10. Glossary

**Purpose:** Define domain terms to ensure shared understanding.

```json
{
  "glossary": [
    {
      "term": "Ingest",
      "definition": "The process of uploading and validating data files",
      "see_also": ["Upload", "Validation"]
    }
  ]
}
```

## Complete Example Structure

```json
{
  "metadata": { ... },
  "mrd_reference": "MRD-2025-001",
  "overview": { ... },
  "functional_requirements": [ ... ],
  "non_functional_requirements": [ ... ],
  "checkpoints": [ ... ],
  "technical_constraints": [ ... ],
  "dependencies": [ ... ],
  "out_of_scope": [ ... ],
  "glossary": [ ... ]
}
```

## Requirement Traceability

Each functional requirement should trace back to an MRD high-level requirement:

```
MRD HLR-001: "Users can upload data files"
  ↓
PRD FR-001: "Upload CSV files up to 100MB"
PRD FR-002: "Upload Excel files up to 50MB"
PRD FR-003: "Display upload progress"
```

This traceability ensures:
- Every PRD requirement serves a business need
- Business requirements aren't lost in translation
- Scope changes can be traced to impact
