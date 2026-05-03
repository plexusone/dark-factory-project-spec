# Implementation Guide

Guidelines for implementing requirements from a PRD.

## Principles

### 1. Minimal Implementation

Implement exactly what the acceptance criteria specify. No more, no less.

**Good:**
```go
// AC: Returns 400 with error message on invalid email
if !isValidEmail(email) {
    return c.JSON(400, map[string]string{"error": "invalid email format"})
}
```

**Bad:**
```go
// Over-engineered: Added features not in spec
if !isValidEmail(email) {
    logger.Warn("invalid email", "email", email, "ip", c.RealIP())
    metrics.Increment("auth.registration.invalid_email")
    notifyAdmin("suspicious registration attempt", email)
    return c.JSON(400, DetailedError{
        Code: "INVALID_EMAIL",
        Message: "invalid email format",
        Suggestions: []string{"Check for typos", "Use a valid domain"},
    })
}
```

### 2. Follow Existing Patterns

Look at how similar things are done in the codebase. Match the style.

```go
// If existing handlers look like this:
func GetUser(c echo.Context) error {
    id := c.Param("id")
    user, err := db.FindUser(id)
    if err != nil {
        return c.JSON(404, err)
    }
    return c.JSON(200, user)
}

// Write new handlers the same way:
func CreateUser(c echo.Context) error {
    var req CreateUserRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(400, err)
    }
    user, err := db.CreateUser(req)
    if err != nil {
        return c.JSON(500, err)
    }
    return c.JSON(201, user)
}
```

### 3. One Test Per Acceptance Criterion

Each acceptance criterion gets its own test:

```go
// AC-001: Given valid email and password, when POST /register, then 201 with user_id
func TestRegister_ValidCredentials_Returns201(t *testing.T) {
    // Given
    body := `{"email": "test@example.com", "password": "SecurePass123"}`

    // When
    resp := POST("/register", body)

    // Then
    assert.Equal(t, 201, resp.StatusCode)
    assert.NotEmpty(t, resp.Body["user_id"])
}

// AC-002: Given invalid email, when POST /register, then 400 with error
func TestRegister_InvalidEmail_Returns400(t *testing.T) {
    // Given
    body := `{"email": "not-an-email", "password": "SecurePass123"}`

    // When
    resp := POST("/register", body)

    // Then
    assert.Equal(t, 400, resp.StatusCode)
    assert.Contains(t, resp.Body["error"], "email")
}
```

### 4. Document Assumptions

When the spec is ambiguous and you make a judgment call, document it:

```go
// NOTE: Spec doesn't specify password hash algorithm.
// Using bcrypt with cost 12 per constitution security standards.
func hashPassword(password string) (string, error) {
    return bcrypt.GenerateFromPassword([]byte(password), 12)
}
```

### 5. Report Discoveries

When you find something unexpected, don't silently handle it. Report it:

```go
// DISCOVERY: Email addresses can contain + for aliasing (user+tag@example.com)
// Current validation allows this, but it could be used to create multiple
// accounts with the same underlying email. Flagging for review.
```

## Implementation Checklist

For each requirement:

- [ ] Read all acceptance criteria
- [ ] Check dependencies are met
- [ ] Review test hints for edge cases
- [ ] Check uncertainty level (high = stop and ask)
- [ ] Write implementation
- [ ] Write test for each acceptance criterion
- [ ] Run tests
- [ ] Document any assumptions
- [ ] Report any discoveries
- [ ] Update implementation record

## Code Quality

### From Constitution

Apply standards from constitution:

```json
{
  "quality": {
    "testing": {
      "unit_test_coverage": "80%"
    }
  },
  "security": {
    "data_protection": {
      "encryption_at_rest": true
    }
  }
}
```

### General Standards

- No hardcoded secrets
- Proper error handling
- Input validation at boundaries
- Logging for debugging (not excessive)
- Comments for non-obvious code only

## Handling Uncertainty

### Low Uncertainty
Proceed with implementation. The spec is clear.

### Medium Uncertainty
Make reasonable assumptions. Document them. Continue.

```go
// ASSUMPTION: "Invalid credentials" error should not distinguish
// between wrong email vs wrong password (security best practice).
```

### High Uncertainty
**STOP.** Do not implement. Return blocked status.

```json
{
  "status": "blocked",
  "blocked_reason": "High uncertainty: spec doesn't define password reset token delivery",
  "discovery": {
    "type": "ambiguity",
    "severity": "high",
    "description": "Password reset needs clarification"
  }
}
```

## Files to Create

Typical requirement implementation creates:

| File Type | Naming | Example |
|-----------|--------|---------|
| Handler | `{feature}.go` | `src/auth/register.go` |
| Tests | `{feature}_test.go` | `src/auth/register_test.go` |
| Types | `types.go` or `{feature}_types.go` | `src/auth/types.go` |
| Migrations | `{number}_{description}.sql` | `migrations/001_users.sql` |

## Commit Strategy

Each requirement can be its own commit:

```
feat(auth): implement user registration (FR-001)

- Add POST /register endpoint
- Validate email format and password strength
- Hash password with bcrypt
- Store user in database
- Return user_id on success

Tests: 5 new tests, all passing
```
