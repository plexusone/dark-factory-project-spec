# MRD Template

This template provides the structure for a Market Requirements Document (MRD). Each section includes guidance on what to include and why it matters.

## 1. Metadata

```json
{
  "metadata": {
    "id": "MRD-YYYY-NNN",
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
- ID format: `MRD-YYYY-NNN` (year + sequential number)
- Version follows semantic versioning
- Status: draft → review → approved → superseded

## 2. Problem Statement

**Purpose:** Clearly articulate the problem being solved. This grounds all subsequent decisions.

```json
{
  "problem_statement": {
    "summary": "One-sentence problem description",
    "current_state": "How things work today",
    "desired_state": "How things should work",
    "impact": "Business impact of the problem (cost, time, risk)"
  }
}
```

**Guidance:**

- **Summary**: Complete the sentence "Users struggle to..." or "The business cannot..."
- **Current state**: Describe the pain in concrete terms. Use specific examples.
- **Desired state**: Describe success. What does "solved" look like?
- **Impact**: Quantify if possible. "Costs X hours/week" or "Risks Y compliance issues"

**Red flags:**
- Vague problems: "Things are slow" → "Search queries take 30+ seconds on datasets >1M rows"
- Solution-as-problem: "We need a cache" → "Response times exceed SLA under load"

## 3. Target Users

**Purpose:** Define who will use this product/feature. Enables prioritization and prevents scope creep.

```json
{
  "target_users": [
    {
      "name": "Persona Name",
      "description": "Who they are, what they do",
      "needs": ["What they need to accomplish"],
      "pain_points": ["Current frustrations"],
      "priority": "primary|secondary|tertiary"
    }
  ]
}
```

**Guidance:**

- List 2-4 user personas
- Primary users are the main audience; design for them first
- Needs should be goals, not features ("upload files" not "drag-and-drop uploader")
- Pain points inform solution design

**Red flags:**
- "Everyone" as a user → Who specifically benefits most?
- Features as needs → What goal does this achieve?

## 4. Success Metrics

**Purpose:** Define how success will be measured. Without metrics, you can't know if you succeeded.

```json
{
  "success_metrics": [
    {
      "id": "SM-001",
      "name": "Metric Name",
      "description": "What is being measured",
      "current_value": "Baseline (if known)",
      "target_value": "Goal",
      "measurement_method": "How to measure",
      "timeframe": "When to evaluate"
    }
  ]
}
```

**Guidance:**

- Include 3-5 metrics
- Mix leading (adoption, usage) and lagging (revenue, satisfaction) indicators
- Make targets specific: "75% satisfaction" not "high satisfaction"
- Describe measurement method to ensure feasibility

**Good metrics:**
- "Reduce average query time from 30s to <5s"
- "Achieve 80% user adoption within 30 days"
- "Reduce support tickets related to X by 50%"

**Red flags:**
- Vanity metrics without business impact
- Unmeasurable targets

## 5. Market Context (Optional)

**Purpose:** Competitive analysis and market positioning.

```json
{
  "market_context": {
    "competitors": [
      {
        "name": "Competitor Name",
        "strengths": ["What they do well"],
        "weaknesses": ["Where they fall short"]
      }
    ],
    "market_trends": ["Relevant trends"],
    "differentiators": ["How this stands out"]
  }
}
```

**Guidance:**

- Skip for internal tools
- For products, identify 2-3 competitors
- Differentiators should be defensible

## 6. High-Level Requirements

**Purpose:** Business-level requirements that will be decomposed into PRD functional requirements.

```json
{
  "high_level_requirements": [
    {
      "id": "HLR-001",
      "description": "What is needed",
      "rationale": "Why this is needed",
      "priority": "must_have|should_have|could_have|wont_have",
      "dependencies": ["Other requirement IDs"]
    }
  ]
}
```

**Guidance:**

- Use MoSCoW prioritization
- Each requirement should be one cohesive capability
- Dependencies help sequence work
- Rationale prevents losing context

**Good requirements:**
- "Users can upload CSV files up to 100MB" (specific, testable)
- "System must comply with GDPR data retention requirements" (clear constraint)

**Bad requirements:**
- "System should be user-friendly" (vague)
- "Handle all file types" (unbounded)

## 7. Constraints

**Purpose:** Document limitations and boundaries. Prevents scope creep and informs trade-offs.

```json
{
  "constraints": [
    {
      "id": "CON-001",
      "type": "technical|business|regulatory|resource|timeline",
      "description": "The constraint",
      "rationale": "Why it exists",
      "negotiable": true|false
    }
  ]
}
```

**Guidance:**

- Non-negotiable constraints are fixed (regulatory, hard deadlines)
- Negotiable constraints can change with trade-offs
- Include rationale so future readers understand context

**Examples:**
- Technical: "Must integrate with existing PostgreSQL database"
- Business: "Budget limited to $50K"
- Regulatory: "Must support SOC 2 audit requirements"
- Timeline: "Must ship before Q3 conference"

## 8. Non-Goals

**Purpose:** Explicitly state what is OUT of scope. Prevents scope creep.

```json
{
  "non_goals": [
    "Feature or capability explicitly excluded",
    "Future consideration not in this version"
  ]
}
```

**Guidance:**

- List 3-5 non-goals
- Include things stakeholders might assume are included
- "Not in V1" is fine if you note it's for later

**Examples:**
- "Mobile app support (future consideration)"
- "Real-time collaboration (single-user editor only)"
- "Integration with Salesforce (out of scope)"

## 9. Assumptions

**Purpose:** Document what you're assuming to be true. Surfaces risks early.

```json
{
  "assumptions": [
    "Assumption 1",
    "Assumption 2"
  ]
}
```

**Guidance:**

- List technical, business, and user assumptions
- If an assumption is wrong, what breaks?
- Consider validating high-risk assumptions before development

**Examples:**
- "Users have stable internet connections"
- "Legal has approved the data sharing approach"
- "The existing API can handle 10x current load"

## 10. Open Questions

**Purpose:** Track unresolved questions. Better to flag than ignore.

```json
{
  "open_questions": [
    {
      "id": "OQ-001",
      "question": "The question",
      "context": "Background info",
      "impact": "What depends on the answer",
      "owner": "Who should answer",
      "due_date": "When needed"
    }
  ]
}
```

**Guidance:**

- Don't block on open questions if you can document uncertainty
- Assign owners and due dates
- Update when answered

## Complete Example Structure

```json
{
  "metadata": { ... },
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
