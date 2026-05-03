# PRD (Product Requirements Document)

The PRD details technical implementation requirements. It answers "how" based on the MRD's "what" and "why".

## Purpose

The PRD:

- Translates MRD requirements to implementable specs
- Defines acceptance criteria for each requirement
- Specifies non-functional requirements
- Establishes checkpoints for validation
- Marks uncertainty explicitly

## Structure

```json
{
  "metadata": { ... },
  "mrd_reference": "MRD-2025-001",
  "overview": { ... },
  "functional_requirements": [ ... ],
  "non_functional_requirements": [ ... ],
  "checkpoints": [ ... ],
  "technical_design": { ... },
  "dependencies": [ ... ],
  "rollout_plan": { ... }
}
```

## Sections

### Functional Requirements

Detailed, implementable requirements:

```json
{
  "functional_requirements": [
    {
      "id": "FR-001",
      "mrd_requirement": "HLR-001",
      "description": "User registration endpoint",
      "acceptance_criteria": [
        "POST /api/auth/register accepts {email, password}",
        "Email must be valid format (RFC 5322)",
        "Password must be at least 8 characters",
        "Password must contain uppercase, lowercase, and number",
        "Email must be unique (409 if exists)",
        "Returns {user_id, email} on success (201)",
        "Returns validation errors on failure (400)"
      ],
      "test_hints": [
        {
          "type": "unit",
          "focus": "Email validation regex"
        },
        {
          "type": "unit",
          "focus": "Password strength validation"
        },
        {
          "type": "integration",
          "focus": "Database uniqueness constraint"
        },
        {
          "type": "edge_case",
          "focus": "Unicode in email local part"
        }
      ],
      "uncertainty": "low",
      "status": "pending",
      "checkpoint": "CP-001"
    }
  ]
}
```

### Acceptance Criteria

Acceptance criteria should be:

- **Specific** - No ambiguity about what "works" means
- **Testable** - Can be verified programmatically
- **Complete** - Cover success, failure, and edge cases

!!! example "Good vs Bad Criteria"
    === "Good"
        - "Returns 401 with `{error: 'invalid_credentials'}` on wrong password"
        - "Locks account after 5 failed attempts within 15 minutes"
        - "JWT expires after 24 hours"

    === "Bad"
        - "Handle authentication errors properly"
        - "Secure login"
        - "Good user experience"

### Uncertainty Markers {#uncertainty}

Requirements can explicitly flag uncertainty:

```json
{
  "id": "FR-003",
  "description": "Handle concurrent uploads",
  "acceptance_criteria": [
    "Multiple users can upload simultaneously"
  ],
  "uncertainty": "high",
  "uncertainty_reason": "Concurrency model not specified in MRD",
  "discovery_prompt": "What happens when two users upload the same file simultaneously? Should we merge, reject, or version?"
}
```

Uncertainty levels:

| Level | Meaning | Action |
|-------|---------|--------|
| `low` | Well understood | Implement directly |
| `medium` | Some unknowns | Clarify during implementation |
| `high` | Significant unknowns | Resolve before implementation |

!!! warning "High uncertainty requires attention"
    High-uncertainty requirements should trigger human review at checkpoints. Don't let AI guess at ambiguous requirements.

### Test Hints

Guidance for test generation:

```json
{
  "test_hints": [
    {
      "type": "unit",
      "focus": "Password hashing uses bcrypt with cost 12"
    },
    {
      "type": "integration",
      "focus": "Token validation against database"
    },
    {
      "type": "edge_case",
      "focus": "Expired token returns 401, not 500"
    },
    {
      "type": "performance",
      "focus": "Login handles 100 concurrent requests"
    },
    {
      "type": "security",
      "focus": "Timing attack resistance on login"
    }
  ]
}
```

Test types:

| Type | Purpose |
|------|---------|
| `unit` | Individual function/component |
| `integration` | Component interaction |
| `edge_case` | Boundary conditions |
| `performance` | Load and latency |
| `security` | Vulnerability testing |

### Non-Functional Requirements

Performance, security, and operational requirements:

```json
{
  "non_functional_requirements": [
    {
      "id": "NFR-001",
      "category": "performance",
      "description": "Login latency under 200ms p99",
      "measurement": "Server-side latency metrics",
      "priority": "must_have"
    },
    {
      "id": "NFR-002",
      "category": "security",
      "description": "Passwords stored with bcrypt cost 12",
      "measurement": "Code review",
      "priority": "must_have"
    },
    {
      "id": "NFR-003",
      "category": "availability",
      "description": "Auth service 99.9% uptime",
      "measurement": "Uptime monitoring",
      "priority": "should_have"
    }
  ]
}
```

Categories:

| Category | Examples |
|----------|----------|
| `performance` | Latency, throughput |
| `security` | Encryption, auth |
| `availability` | Uptime, redundancy |
| `scalability` | Load handling |
| `maintainability` | Code quality |

### Requirement Status

Track implementation progress:

```json
{
  "id": "FR-001",
  "status": "implemented"
}
```

Status values:

| Status | Meaning |
|--------|---------|
| `pending` | Not started |
| `in_progress` | Being implemented |
| `blocked` | Waiting on something |
| `implemented` | Code complete |
| `validated` | Tested and verified |

## Checkpoints

See [Checkpoints](checkpoints.md) for details on defining validation gates.

## Best Practices

1. **Link to MRD** - Every FR should trace to an HLR
2. **Be specific** - "200ms p99" not "fast"
3. **Include failure cases** - What happens when things go wrong?
4. **Mark uncertainty** - Don't hide what you don't know
5. **Add test hints** - Guide comprehensive test coverage
6. **Define checkpoints** - Create natural validation points
