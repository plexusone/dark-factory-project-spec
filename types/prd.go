package types

// PRD represents a Product Requirements Document.
// It contains detailed functional requirements derived from an MRD.
type PRD struct {
	// Metadata about the document itself
	Metadata DocumentMetadata `json:"metadata"`

	// MRDReference links to the parent MRD
	MRDReference string `json:"mrd_reference"`

	// Overview provides context for the PRD
	Overview PRDOverview `json:"overview"`

	// FunctionalRequirements are the detailed requirements
	FunctionalRequirements []FunctionalRequirement `json:"functional_requirements"`

	// NonFunctionalRequirements cover quality attributes
	NonFunctionalRequirements []NonFunctionalRequirement `json:"non_functional_requirements,omitempty"`

	// Checkpoints define validation points during execution
	Checkpoints []Checkpoint `json:"checkpoints,omitempty"`

	// TechnicalConstraints are implementation constraints
	TechnicalConstraints []TechnicalConstraint `json:"technical_constraints,omitempty"`

	// Dependencies lists external dependencies
	Dependencies []Dependency `json:"dependencies,omitempty"`

	// OutOfScope explicitly states what is not included
	OutOfScope []string `json:"out_of_scope,omitempty"`

	// Glossary defines domain terms
	Glossary []GlossaryTerm `json:"glossary,omitempty"`
}

// PRDOverview provides context for the PRD.
type PRDOverview struct {
	// Summary is a brief description of what is being built
	Summary string `json:"summary"`

	// Goals lists the goals this PRD addresses
	Goals []string `json:"goals"`

	// TargetRelease is the intended release version
	TargetRelease string `json:"target_release,omitempty"`
}

// FunctionalRequirement is a detailed requirement.
type FunctionalRequirement struct {
	// ID uniquely identifies this requirement (e.g., "FR-001")
	ID string `json:"id"`

	// MRDRequirement links to the parent MRD requirement
	MRDRequirement string `json:"mrd_requirement,omitempty"`

	// Title is a short name for the requirement
	Title string `json:"title"`

	// Description explains what is required
	Description string `json:"description"`

	// AcceptanceCriteria define when the requirement is met
	AcceptanceCriteria []AcceptanceCriterion `json:"acceptance_criteria"`

	// TestHints guide test generation
	TestHints []TestHint `json:"test_hints,omitempty"`

	// Priority indicates importance
	Priority RequirementPriority `json:"priority"`

	// Status tracks implementation progress
	Status RequirementStatus `json:"status"`

	// BlockedReason explains why the requirement is blocked
	BlockedReason string `json:"blocked_reason,omitempty"`

	// Dependencies lists other requirement IDs this depends on
	Dependencies []string `json:"dependencies,omitempty"`

	// Checkpoint indicates when to validate this requirement
	Checkpoint string `json:"checkpoint,omitempty"`

	// Uncertainty indicates confidence level
	Uncertainty UncertaintyLevel `json:"uncertainty"`

	// UncertaintyReason explains why uncertainty is high
	UncertaintyReason string `json:"uncertainty_reason,omitempty"`

	// DiscoveryPrompt suggests questions to clarify uncertainty
	DiscoveryPrompt string `json:"discovery_prompt,omitempty"`
}

// AcceptanceCriterion defines when a requirement is met.
type AcceptanceCriterion struct {
	// ID uniquely identifies this criterion
	ID string `json:"id"`

	// Given is the precondition
	Given string `json:"given"`

	// When is the action
	When string `json:"when"`

	// Then is the expected outcome
	Then string `json:"then"`
}

// TestHint guides test generation.
type TestHint struct {
	// Type categorizes the test
	Type TestType `json:"type"`

	// Focus describes what to test
	Focus string `json:"focus"`

	// Inputs suggests test inputs
	Inputs []string `json:"inputs,omitempty"`

	// ExpectedBehaviors describes expected outcomes
	ExpectedBehaviors []string `json:"expected_behaviors,omitempty"`
}

// TestType categorizes tests.
type TestType string

const (
	// TestUnit is a unit test
	TestUnit TestType = "unit"

	// TestIntegration is an integration test
	TestIntegration TestType = "integration"

	// TestE2E is an end-to-end test
	TestE2E TestType = "e2e"

	// TestEdgeCase is an edge case test
	TestEdgeCase TestType = "edge_case"

	// TestPerformance is a performance test
	TestPerformance TestType = "performance"

	// TestSecurity is a security test
	TestSecurity TestType = "security"
)

// RequirementStatus tracks implementation progress.
type RequirementStatus string

const (
	// StatusPending means not started
	StatusPending RequirementStatus = "pending"

	// StatusInProgress means being worked on
	StatusInProgress RequirementStatus = "in_progress"

	// StatusBlocked means waiting on something
	StatusBlocked RequirementStatus = "blocked"

	// StatusImplemented means code is written
	StatusImplemented RequirementStatus = "implemented"

	// StatusValidated means tests pass
	StatusValidated RequirementStatus = "validated"
)

// UncertaintyLevel indicates confidence in the requirement.
type UncertaintyLevel string

const (
	// UncertaintyLow means high confidence
	UncertaintyLow UncertaintyLevel = "low"

	// UncertaintyMedium means some unknowns
	UncertaintyMedium UncertaintyLevel = "medium"

	// UncertaintyHigh means significant unknowns
	UncertaintyHigh UncertaintyLevel = "high"
)

// NonFunctionalRequirement covers quality attributes.
type NonFunctionalRequirement struct {
	// ID uniquely identifies this requirement
	ID string `json:"id"`

	// Category is the type of NFR
	Category NFRCategory `json:"category"`

	// Description explains the requirement
	Description string `json:"description"`

	// Measure is how to measure compliance
	Measure string `json:"measure"`

	// Target is the specific target to meet
	Target string `json:"target"`

	// Priority indicates importance
	Priority RequirementPriority `json:"priority"`
}

// NFRCategory categorizes non-functional requirements.
type NFRCategory string

const (
	// NFRPerformance covers speed and throughput
	NFRPerformance NFRCategory = "performance"

	// NFRScalability covers growth handling
	NFRScalability NFRCategory = "scalability"

	// NFRReliability covers uptime and recovery
	NFRReliability NFRCategory = "reliability"

	// NFRSecurity covers protection
	NFRSecurity NFRCategory = "security"

	// NFRUsability covers user experience
	NFRUsability NFRCategory = "usability"

	// NFRMaintainability covers code quality
	NFRMaintainability NFRCategory = "maintainability"

	// NFRAccessibility covers accessibility standards
	NFRAccessibility NFRCategory = "accessibility"
)

// TechnicalConstraint is an implementation constraint.
type TechnicalConstraint struct {
	// ID uniquely identifies this constraint
	ID string `json:"id"`

	// Area is the technical area affected
	Area string `json:"area"`

	// Constraint describes the limitation
	Constraint string `json:"constraint"`

	// Rationale explains why
	Rationale string `json:"rationale"`
}

// Dependency lists external dependencies.
type Dependency struct {
	// ID uniquely identifies this dependency
	ID string `json:"id"`

	// Name is the dependency name
	Name string `json:"name"`

	// Type categorizes the dependency
	Type DependencyType `json:"type"`

	// Description explains the dependency
	Description string `json:"description"`

	// Version is the required version (if applicable)
	Version string `json:"version,omitempty"`

	// Owner is who is responsible
	Owner string `json:"owner,omitempty"`

	// Status is the current state
	Status DependencyStatus `json:"status"`
}

// DependencyType categorizes dependencies.
type DependencyType string

const (
	// DependencyService is an external service
	DependencyService DependencyType = "service"

	// DependencyLibrary is a software library
	DependencyLibrary DependencyType = "library"

	// DependencyAPI is an external API
	DependencyAPI DependencyType = "api"

	// DependencyTeam is another team's deliverable
	DependencyTeam DependencyType = "team"

	// DependencyData is a data source
	DependencyData DependencyType = "data"
)

// DependencyStatus is the current state of a dependency.
type DependencyStatus string

const (
	// DependencyAvailable means ready to use
	DependencyAvailable DependencyStatus = "available"

	// DependencyInProgress means being worked on
	DependencyInProgress DependencyStatus = "in_progress"

	// DependencyBlocked means not available
	DependencyBlocked DependencyStatus = "blocked"

	// DependencyUnknown means status unclear
	DependencyUnknown DependencyStatus = "unknown"
)

// GlossaryTerm defines a domain term.
type GlossaryTerm struct {
	// Term is the word or phrase
	Term string `json:"term"`

	// Definition explains the meaning
	Definition string `json:"definition"`

	// SeeAlso links to related terms
	SeeAlso []string `json:"see_also,omitempty"`
}
