# Section Questions

This document contains probing questions for each section of MRD and PRD documents. Use these questions during the interview phase to ensure completeness.

## MRD Sections

### Problem Statement

**Core questions:**
- What problem are you trying to solve?
- Who experiences this problem?
- How do people work around it today?
- What happens if we don't solve this?

**Probing questions:**
- Can you give me a specific example of this problem occurring?
- How often does this problem occur? Daily? Weekly?
- What's the cost of this problem? (time, money, risk, frustration)
- Is this a new problem or an existing one getting worse?
- Have you tried to solve this before? What happened?

**Red flag questions:**
- Are you describing a problem or a solution? (If solution, ask "What problem does that solve?")
- Is this problem urgent or important? (Helps prioritize)

### Target Users

**Core questions:**
- Who will use this? Be specific.
- What do they do in their job/life?
- What are they trying to accomplish?
- How technically sophisticated are they?

**Probing questions:**
- Of all the potential users, who benefits most? (primary user)
- Are there users who interact with this indirectly?
- Do different users have conflicting needs?
- What tools/products do they use today?
- What frustrates them about current solutions?

**Edge case questions:**
- What about new users vs. power users?
- What about users with accessibility needs?
- What about users in different regions/languages?

### Success Metrics

**Core questions:**
- How will you know this succeeded?
- What does "good enough" look like?
- What's the baseline today?
- When will you measure success?

**Probing questions:**
- Can you actually measure this? How?
- What's the minimum improvement that would be meaningful?
- Are there leading indicators we can track early?
- What would make this a failure?
- Are there vanity metrics we should avoid?

**Reality check questions:**
- If you hit this metric but users hate it, does it matter?
- What other metrics might move as a side effect?

### High-Level Requirements

**Core questions:**
- What capabilities are needed to solve the problem?
- What's essential vs. nice-to-have?
- What's the smallest thing that would be useful?

**Prioritization questions:**
- If you could only have one of these, which one?
- What's blocking adoption today? (must-haves)
- What would delight users but isn't essential? (should-haves)
- What are you explicitly leaving for later? (won't-haves)

**Dependency questions:**
- Do any requirements depend on others?
- Are there prerequisites outside this scope?
- Are there integration points with other systems?

### Constraints

**Core questions:**
- What can't we change? (technical, business, regulatory)
- What's the timeline?
- What's the budget/resource limit?
- Are there compliance requirements?

**Probing questions:**
- Which constraints are truly fixed vs. negotiable?
- Why does this constraint exist?
- What happens if we violate this constraint?
- Are there constraints you haven't mentioned because they seem obvious?

**Hidden constraint questions:**
- Are there existing systems we must integrate with?
- Are there technologies we must or cannot use?
- Are there organizational politics affecting this?

### Non-Goals

**Core questions:**
- What are you explicitly NOT building?
- What might people assume is included but isn't?
- What's being deferred to future versions?

**Probing questions:**
- If a stakeholder asked for X, would you add it?
- Are there features competitors have that you're intentionally skipping?
- What would make this too complex for V1?

### Assumptions

**Core questions:**
- What are you assuming to be true?
- What would break if these assumptions are wrong?
- How confident are you in each assumption?

**Technical assumptions:**
- What's assumed about the existing infrastructure?
- What's assumed about data quality/availability?
- What's assumed about scale/load?

**Business assumptions:**
- What's assumed about user behavior?
- What's assumed about adoption rate?
- What's assumed about stakeholder support?

### Open Questions

**Core questions:**
- What don't you know yet?
- What decisions are pending?
- Who needs to provide input?

**Probing questions:**
- What's the riskiest unknown?
- Can we proceed without knowing this, or is it blocking?
- Who can answer this question?
- What's the deadline for resolving this?

---

## PRD Sections

### Functional Requirements

**Core questions:**
- What exactly should the system do?
- How should it behave in the normal case?
- How should it behave in error cases?

**Completeness questions:**
- What happens when input is invalid?
- What happens when the system fails?
- What happens at boundaries? (empty, max, zero)
- What happens with concurrent access?
- What happens with different permission levels?

**Testability questions:**
- How will we know this requirement is met?
- What's the acceptance criteria?
- Can you describe a test case for this?

**Edge case questions:**
- What's the maximum/minimum expected?
- What if the user does something unexpected?
- What about internationalization?
- What about accessibility?

### Acceptance Criteria

**Core questions:**
- What's the precondition (Given)?
- What action triggers this (When)?
- What's the expected outcome (Then)?

**Completeness questions:**
- Have we covered the happy path?
- Have we covered error paths?
- Have we covered edge cases?
- Have we covered permission scenarios?

**Quality questions:**
- Is this criteria testable as written?
- Is this criteria specific enough?
- Would two people interpret this the same way?

### Test Hints

**Core questions:**
- What should unit tests cover?
- What integration points need testing?
- What edge cases are most risky?

**Coverage questions:**
- What would a security test check?
- What would a performance test check?
- What about data validation?
- What about error handling?

### Checkpoints

**Core questions:**
- Where should we pause for validation?
- What requirements are foundational?
- What changes are irreversible?

**Timing questions:**
- What needs to be validated before moving on?
- Are there dependencies between requirements?
- Where do integration risks exist?

### Dependencies

**Core questions:**
- What external services/APIs are needed?
- What libraries/frameworks are required?
- What other teams' work is this blocked by?

**Status questions:**
- Is this dependency available now?
- Who owns this dependency?
- What's the risk if this dependency is delayed?

### Uncertainty

**Core questions:**
- What are you unsure about in this requirement?
- Why is there uncertainty?
- What question would resolve this?

**Risk questions:**
- What's the cost of guessing wrong?
- Can we proceed with uncertainty, or is it blocking?
- Should this have a checkpoint for validation?

---

## General Probing Techniques

### Clarification
- "Can you give me an example?"
- "What do you mean by [term]?"
- "How does that differ from [alternative]?"

### Completeness
- "What else should I know?"
- "Is there anything I haven't asked about?"
- "What would someone else working on this need to know?"

### Edge Cases
- "What happens if [boundary condition]?"
- "What about the negative case?"
- "What if this fails?"

### Priority
- "Is this essential or nice-to-have?"
- "What would you cut if you had to?"
- "What's the minimum viable version?"

### Assumptions
- "Are you assuming [X]?"
- "Would this change if [condition]?"
- "What are you taking for granted here?"
