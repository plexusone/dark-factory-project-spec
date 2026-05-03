# MRD (Market Requirements Document)

The MRD captures high-level business requirements and market context. It answers "what" and "why" without diving into "how".

## Purpose

The MRD:

- Defines the problem being solved
- Identifies target users and their needs
- Establishes success metrics
- Lists high-level requirements
- Documents constraints and non-goals

## Structure

```json
{
  "metadata": { ... },
  "constitution_ref": "specs/constitution.json",
  "problem_statement": { ... },
  "target_users": [ ... ],
  "success_metrics": [ ... ],
  "market_context": { ... },
  "high_level_requirements": [ ... ],
  "constraints": [ ... ],
  "non_goals": [ ... ],
  "assumptions": [ ... ],
  "open_questions": [ ... ]
}
```

## Sections

### Metadata

Document identification and versioning:

```json
{
  "metadata": {
    "id": "MRD-2025-001",
    "title": "User Authentication",
    "feature": "user-authentication",
    "version": "1.0.0",
    "status": "approved",
    "authors": ["Product Manager"],
    "created_at": "2025-01-15T00:00:00Z",
    "updated_at": "2025-01-15T00:00:00Z",
    "approved_at": "2025-01-20T00:00:00Z",
    "approved_by": ["Engineering Lead", "CTO"]
  }
}
```

Status values:

| Status | Meaning |
|--------|---------|
| `draft` | Being written |
| `review` | Under review |
| `approved` | Ready for PRD/implementation |
| `superseded` | Replaced by newer version |

### Problem Statement

Clearly defines the problem:

```json
{
  "problem_statement": {
    "summary": "Users cannot securely access protected resources",
    "current_state": "No authentication exists. Users access all resources anonymously.",
    "desired_state": "Users can register, login, and access resources based on permissions.",
    "impact": "Without auth, we cannot implement any user-specific features or monetization."
  }
}
```

!!! tip "Writing good problem statements"
    - Be specific about the current pain
    - Quantify impact where possible
    - Avoid solutioning in the problem statement

### Target Users

Defines user personas:

```json
{
  "target_users": [
    {
      "name": "End User",
      "description": "Individual using the application",
      "needs": [
        "Quick registration",
        "Secure login",
        "Password recovery"
      ],
      "pain_points": [
        "Remembering passwords",
        "Complex registration forms"
      ],
      "priority": "primary"
    },
    {
      "name": "Administrator",
      "description": "Internal user managing the system",
      "needs": [
        "User management",
        "Audit logs"
      ],
      "priority": "secondary"
    }
  ]
}
```

Priority levels:

| Priority | Meaning |
|----------|---------|
| `primary` | Main target users |
| `secondary` | Important but not main |
| `tertiary` | Nice to support |

### Success Metrics

How success is measured:

```json
{
  "success_metrics": [
    {
      "id": "SM-001",
      "name": "Registration Completion Rate",
      "description": "Percentage of users who complete registration",
      "current_value": "N/A",
      "target_value": ">90%",
      "measurement_method": "Analytics: registrations started vs completed",
      "timeframe": "30 days post-launch"
    },
    {
      "id": "SM-002",
      "name": "Login Success Rate",
      "description": "Percentage of login attempts that succeed",
      "target_value": ">99%",
      "measurement_method": "Server logs: 200 vs 401 responses"
    }
  ]
}
```

### High-Level Requirements

Business-level requirements:

```json
{
  "high_level_requirements": [
    {
      "id": "HLR-001",
      "description": "Users can register with email and password",
      "rationale": "Foundation for user identity",
      "priority": "must_have",
      "dependencies": []
    },
    {
      "id": "HLR-002",
      "description": "Users can log in with credentials",
      "rationale": "Access to protected resources",
      "priority": "must_have",
      "dependencies": ["HLR-001"]
    },
    {
      "id": "HLR-003",
      "description": "Users can reset forgotten password",
      "rationale": "Self-service reduces support burden",
      "priority": "should_have",
      "dependencies": ["HLR-001"]
    }
  ]
}
```

Priority levels (MoSCoW):

| Priority | Meaning |
|----------|---------|
| `must_have` | Essential for launch |
| `should_have` | Important but not critical |
| `could_have` | Nice to have |
| `wont_have` | Explicitly excluded |

### Constraints

Limitations and boundaries:

```json
{
  "constraints": [
    {
      "id": "C-001",
      "type": "technical",
      "description": "Must use existing PostgreSQL database",
      "rationale": "No budget for new infrastructure",
      "negotiable": false
    },
    {
      "id": "C-002",
      "type": "timeline",
      "description": "Must launch by Q2 2025",
      "rationale": "Contractual obligation",
      "negotiable": false
    },
    {
      "id": "C-003",
      "type": "regulatory",
      "description": "Must comply with GDPR",
      "rationale": "EU users",
      "negotiable": false
    }
  ]
}
```

Constraint types:

| Type | Examples |
|------|----------|
| `technical` | Existing systems, integrations |
| `business` | Budget, resources |
| `regulatory` | Compliance, legal |
| `resource` | Team size, skills |
| `timeline` | Deadlines |

### Non-Goals

What is explicitly out of scope:

```json
{
  "non_goals": [
    "Social login (OAuth) - defer to v2",
    "Multi-factor authentication - defer to v2",
    "Single sign-on (SSO) - enterprise feature",
    "Biometric authentication - mobile-only concern"
  ]
}
```

!!! warning "Non-goals are critical"
    Explicitly stating what you're NOT doing prevents scope creep and misaligned expectations.

### Open Questions

Unresolved questions:

```json
{
  "open_questions": [
    {
      "id": "OQ-001",
      "question": "Should we support username in addition to email?",
      "context": "Some users prefer usernames for privacy",
      "impact": "Affects registration flow and login options",
      "owner": "Product Manager",
      "due_date": "2025-01-25T00:00:00Z"
    }
  ]
}
```

## Best Practices

1. **Stay high-level** - Save implementation details for PRD
2. **Be specific** - Vague requirements lead to wrong implementations
3. **Include rationale** - Explain WHY, not just WHAT
4. **List non-goals explicitly** - Prevents scope creep
5. **Track open questions** - Don't hide uncertainty
6. **Define success** - Metrics make success measurable
