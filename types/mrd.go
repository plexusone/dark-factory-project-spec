// Package types defines the core data structures for MRD/PRD specifications.
package types

// MRD represents a Market Requirements Document.
// It captures the high-level business requirements and market context
// for a product or feature.
type MRD struct {
	// Metadata about the document itself
	Metadata DocumentMetadata `json:"metadata"`

	// ConstitutionRef references the organization constitution this inherits from
	ConstitutionRef string `json:"constitution_ref,omitempty"`

	// ConstitutionOverrides lists any deviations from the constitution
	ConstitutionOverrides *ConstitutionOverride `json:"constitution_overrides,omitempty"`

	// ProblemStatement describes the problem being solved
	ProblemStatement ProblemStatement `json:"problem_statement"`

	// TargetUsers describes who will use this product/feature
	TargetUsers []TargetUser `json:"target_users"`

	// SuccessMetrics defines how success will be measured
	SuccessMetrics []SuccessMetric `json:"success_metrics"`

	// MarketContext provides competitive and market analysis
	MarketContext MarketContext `json:"market_context,omitempty"`

	// HighLevelRequirements are business-level requirements
	HighLevelRequirements []HighLevelRequirement `json:"high_level_requirements"`

	// Constraints are limitations or boundaries
	Constraints []Constraint `json:"constraints,omitempty"`

	// NonGoals explicitly state what is out of scope
	NonGoals []string `json:"non_goals,omitempty"`

	// Assumptions document what we're assuming to be true
	Assumptions []string `json:"assumptions,omitempty"`

	// OpenQuestions tracks unresolved questions
	OpenQuestions []OpenQuestion `json:"open_questions,omitempty"`
}

// DocumentMetadata contains metadata about the spec document.
type DocumentMetadata struct {
	// ID is a unique identifier for this document
	ID string `json:"id"`

	// Title is the human-readable title
	Title string `json:"title"`

	// Feature is the kebab-case feature name this spec belongs to
	// Used for organizing docs when strategy is "feature"
	// e.g., "user-authentication", "data-ingestion"
	Feature string `json:"feature,omitempty"`

	// Version follows semantic versioning
	Version string `json:"version"`

	// PreviousVersion is the version this was derived from
	PreviousVersion string `json:"previous_version,omitempty"`

	// Status is the document lifecycle status
	Status DocumentStatus `json:"status"`

	// Authors lists the document authors
	Authors []string `json:"authors"`

	// CreatedAt is when the document was created (RFC3339)
	CreatedAt string `json:"created_at"`

	// UpdatedAt is when the document was last updated (RFC3339)
	UpdatedAt string `json:"updated_at"`

	// ApprovedAt is when the document was approved (RFC3339)
	ApprovedAt string `json:"approved_at,omitempty"`

	// ApprovedBy lists who approved the document
	ApprovedBy []string `json:"approved_by,omitempty"`

	// Amendments track changes after approval
	Amendments []Amendment `json:"amendments,omitempty"`
}

// DocumentStatus represents the lifecycle status of a document.
type DocumentStatus string

const (
	// StatusDraft means the document is being written
	StatusDraft DocumentStatus = "draft"

	// StatusReview means the document is under review
	StatusReview DocumentStatus = "review"

	// StatusApproved means the document has been approved
	StatusApproved DocumentStatus = "approved"

	// StatusSuperseded means the document has been replaced
	StatusSuperseded DocumentStatus = "superseded"
)

// Amendment tracks a change to an approved document.
type Amendment struct {
	// ID uniquely identifies this amendment
	ID string `json:"id"`

	// Description explains what was changed
	Description string `json:"description"`

	// Reason explains why the change was made
	Reason string `json:"reason"`

	// ChangedAt is when the amendment was made (RFC3339)
	ChangedAt string `json:"changed_at"`

	// ChangedBy is who made the amendment
	ChangedBy string `json:"changed_by"`

	// AffectedRequirements lists requirement IDs that changed
	AffectedRequirements []string `json:"affected_requirements,omitempty"`
}

// ProblemStatement describes the problem being solved.
type ProblemStatement struct {
	// Summary is a brief description of the problem
	Summary string `json:"summary"`

	// CurrentState describes how things work today
	CurrentState string `json:"current_state"`

	// DesiredState describes how things should work
	DesiredState string `json:"desired_state"`

	// Impact describes the business impact of the problem
	Impact string `json:"impact"`
}

// TargetUser describes a user persona.
type TargetUser struct {
	// Name is the persona name (e.g., "Developer", "Data Analyst")
	Name string `json:"name"`

	// Description explains who this user is
	Description string `json:"description"`

	// Needs lists what this user needs
	Needs []string `json:"needs"`

	// PainPoints lists current frustrations
	PainPoints []string `json:"pain_points,omitempty"`

	// Priority indicates how important this user segment is
	Priority UserPriority `json:"priority"`
}

// UserPriority indicates the priority of a user segment.
type UserPriority string

const (
	// PriorityPrimary is the main target user
	PriorityPrimary UserPriority = "primary"

	// PrioritySecondary is an important but not main user
	PrioritySecondary UserPriority = "secondary"

	// PriorityTertiary is a nice-to-have user segment
	PriorityTertiary UserPriority = "tertiary"
)

// SuccessMetric defines how success is measured.
type SuccessMetric struct {
	// ID uniquely identifies this metric
	ID string `json:"id"`

	// Name is the metric name
	Name string `json:"name"`

	// Description explains what is being measured
	Description string `json:"description"`

	// CurrentValue is the baseline value
	CurrentValue string `json:"current_value,omitempty"`

	// TargetValue is the goal
	TargetValue string `json:"target_value"`

	// MeasurementMethod explains how to measure
	MeasurementMethod string `json:"measurement_method"`

	// Timeframe is when the target should be achieved
	Timeframe string `json:"timeframe,omitempty"`
}

// MarketContext provides competitive and market analysis.
type MarketContext struct {
	// Competitors lists relevant competitors
	Competitors []Competitor `json:"competitors,omitempty"`

	// MarketTrends describes relevant market trends
	MarketTrends []string `json:"market_trends,omitempty"`

	// Differentiators explain how this stands out
	Differentiators []string `json:"differentiators,omitempty"`
}

// Competitor describes a competitive product.
type Competitor struct {
	// Name is the competitor name
	Name string `json:"name"`

	// Strengths lists their advantages
	Strengths []string `json:"strengths,omitempty"`

	// Weaknesses lists their disadvantages
	Weaknesses []string `json:"weaknesses,omitempty"`
}

// HighLevelRequirement is a business-level requirement.
type HighLevelRequirement struct {
	// ID uniquely identifies this requirement
	ID string `json:"id"`

	// Description explains what is needed
	Description string `json:"description"`

	// Rationale explains why this is needed
	Rationale string `json:"rationale"`

	// Priority indicates importance
	Priority RequirementPriority `json:"priority"`

	// Dependencies lists other requirement IDs this depends on
	Dependencies []string `json:"dependencies,omitempty"`
}

// RequirementPriority indicates the priority of a requirement.
type RequirementPriority string

const (
	// PriorityMustHave is essential for launch
	PriorityMustHave RequirementPriority = "must_have"

	// PriorityShouldHave is important but not critical
	PriorityShouldHave RequirementPriority = "should_have"

	// PriorityCouldHave is nice to have
	PriorityCouldHave RequirementPriority = "could_have"

	// PriorityWontHave is explicitly excluded
	PriorityWontHave RequirementPriority = "wont_have"
)

// Constraint is a limitation or boundary.
type Constraint struct {
	// ID uniquely identifies this constraint
	ID string `json:"id"`

	// Type is the kind of constraint
	Type ConstraintType `json:"type"`

	// Description explains the constraint
	Description string `json:"description"`

	// Rationale explains why this constraint exists
	Rationale string `json:"rationale,omitempty"`

	// Negotiable indicates if this can be changed
	Negotiable bool `json:"negotiable"`
}

// ConstraintType categorizes constraints.
type ConstraintType string

const (
	// ConstraintTechnical is a technical limitation
	ConstraintTechnical ConstraintType = "technical"

	// ConstraintBusiness is a business limitation
	ConstraintBusiness ConstraintType = "business"

	// ConstraintRegulatory is a legal/compliance requirement
	ConstraintRegulatory ConstraintType = "regulatory"

	// ConstraintResource is a resource limitation
	ConstraintResource ConstraintType = "resource"

	// ConstraintTimeline is a time-based constraint
	ConstraintTimeline ConstraintType = "timeline"
)

// OpenQuestion tracks an unresolved question.
type OpenQuestion struct {
	// ID uniquely identifies this question
	ID string `json:"id"`

	// Question is the question itself
	Question string `json:"question"`

	// Context provides background
	Context string `json:"context,omitempty"`

	// Impact describes what depends on the answer
	Impact string `json:"impact,omitempty"`

	// Owner is who is responsible for answering
	Owner string `json:"owner,omitempty"`

	// DueDate is when an answer is needed (RFC3339)
	DueDate string `json:"due_date,omitempty"`

	// Answer is the resolved answer
	Answer string `json:"answer,omitempty"`

	// AnsweredAt is when it was answered (RFC3339)
	AnsweredAt string `json:"answered_at,omitempty"`
}
