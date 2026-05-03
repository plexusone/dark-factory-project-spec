package types

// SpecEvaluation represents the evaluation of a spec document.
// It follows the structured-evaluation format for compatibility.
type SpecEvaluation struct {
	// Metadata about the evaluation
	Metadata EvaluationMetadata `json:"metadata"`

	// Subject identifies what was evaluated
	Subject EvaluationSubject `json:"subject"`

	// Categories contains scores for each evaluation category
	Categories []CategoryScore `json:"categories"`

	// Findings contains specific issues found
	Findings []Finding `json:"findings"`

	// Summary provides overall assessment
	Summary EvaluationSummary `json:"summary"`

	// Decision is the GO/NO-GO recommendation
	Decision Decision `json:"decision"`
}

// EvaluationMetadata contains metadata about the evaluation.
type EvaluationMetadata struct {
	// EvaluationID uniquely identifies this evaluation
	EvaluationID string `json:"evaluation_id"`

	// EvaluatedAt is when the evaluation was performed (RFC3339)
	EvaluatedAt string `json:"evaluated_at"`

	// Evaluator identifies who/what performed the evaluation
	Evaluator string `json:"evaluator"`

	// RubricVersion identifies the rubric used
	RubricVersion string `json:"rubric_version"`

	// ModelID identifies the model used for evaluation
	ModelID string `json:"model_id,omitempty"`

	// JudgeRuns captures multi-judge aggregation (if used)
	JudgeRuns []JudgeRun `json:"judge_runs,omitempty"`
}

// JudgeRun captures a single judge's evaluation for aggregation.
type JudgeRun struct {
	// JudgeID identifies this judge
	JudgeID string `json:"judge_id"`

	// ModelID identifies the model used
	ModelID string `json:"model_id"`

	// Temperature is the sampling temperature used
	Temperature float64 `json:"temperature"`

	// Decision is this judge's decision
	Decision DecisionType `json:"decision"`

	// CategoryScores are this judge's category scores
	CategoryScores []CategoryScore `json:"category_scores"`

	// ExecutedAt is when this judge ran (RFC3339)
	ExecutedAt string `json:"executed_at"`
}

// EvaluationSubject identifies what was evaluated.
type EvaluationSubject struct {
	// Type is the document type (mrd, prd)
	Type SpecType `json:"type"`

	// ID is the document ID
	ID string `json:"id"`

	// Version is the document version
	Version string `json:"version"`

	// Title is the document title
	Title string `json:"title"`
}

// SpecType identifies the type of spec document.
type SpecType string

const (
	// SpecTypeMRD is a Market Requirements Document
	SpecTypeMRD SpecType = "mrd"

	// SpecTypePRD is a Product Requirements Document
	SpecTypePRD SpecType = "prd"
)

// CategoryScore represents a score for an evaluation category.
type CategoryScore struct {
	// Category is the category name
	Category string `json:"category"`

	// Weight is the category weight (0-1)
	Weight float64 `json:"weight"`

	// Score is the score (0-10)
	Score float64 `json:"score"`

	// Rationale explains the score
	Rationale string `json:"rationale"`
}

// Finding represents a specific issue found during evaluation.
type Finding struct {
	// ID uniquely identifies this finding
	ID string `json:"id"`

	// Severity indicates the importance of this finding
	Severity FindingSeverity `json:"severity"`

	// Category is the evaluation category this relates to
	Category string `json:"category"`

	// Title is a short description
	Title string `json:"title"`

	// Description provides full details
	Description string `json:"description"`

	// Location points to the relevant section/requirement
	Location string `json:"location,omitempty"`

	// Suggestion provides guidance on fixing
	Suggestion string `json:"suggestion,omitempty"`

	// Blocking indicates if this prevents approval
	Blocking bool `json:"blocking"`
}

// FindingSeverity indicates the importance of a finding.
type FindingSeverity string

const (
	// SeverityCriticalFinding is a critical issue that must be fixed
	SeverityCriticalFinding FindingSeverity = "critical"

	// SeverityHighFinding is a high-priority issue that must be fixed
	SeverityHighFinding FindingSeverity = "high"

	// SeverityMediumFinding requires attention but doesn't block
	SeverityMediumFinding FindingSeverity = "medium"

	// SeverityLowFinding is minor and can be addressed later
	SeverityLowFinding FindingSeverity = "low"

	// SeverityInfoFinding is informational only
	SeverityInfoFinding FindingSeverity = "info"
)

// EvaluationSummary provides overall assessment.
type EvaluationSummary struct {
	// OverallScore is the weighted average score (0-10)
	OverallScore float64 `json:"overall_score"`

	// Strengths lists what is done well
	Strengths []string `json:"strengths"`

	// Weaknesses lists areas needing improvement
	Weaknesses []string `json:"weaknesses"`

	// FindingCounts summarizes findings by severity
	FindingCounts FindingCounts `json:"finding_counts"`
}

// FindingCounts summarizes findings by severity.
type FindingCounts struct {
	// Critical is the count of critical findings
	Critical int `json:"critical"`

	// High is the count of high findings
	High int `json:"high"`

	// Medium is the count of medium findings
	Medium int `json:"medium"`

	// Low is the count of low findings
	Low int `json:"low"`

	// Info is the count of info findings
	Info int `json:"info"`
}

// Decision is the GO/NO-GO recommendation.
type Decision struct {
	// Type is the decision type
	Type DecisionType `json:"type"`

	// Rationale explains the decision
	Rationale string `json:"rationale"`

	// Conditions lists any conditions for approval
	Conditions []string `json:"conditions,omitempty"`

	// RequiresHumanReview indicates if human review is needed
	RequiresHumanReview bool `json:"requires_human_review"`

	// HumanReviewReason explains why human review is needed
	HumanReviewReason string `json:"human_review_reason,omitempty"`
}

// DecisionType is the type of GO/NO-GO decision.
type DecisionType string

const (
	// DecisionGo means approved for development
	DecisionGo DecisionType = "go"

	// DecisionNoGo means not approved for development
	DecisionNoGo DecisionType = "no_go"

	// DecisionConditional means approved with conditions
	DecisionConditional DecisionType = "conditional"
)

// EvaluationRubric defines the rubric for evaluating specs.
type EvaluationRubric struct {
	// Version identifies this rubric version
	Version string `json:"version"`

	// SpecType is the type of spec this rubric evaluates
	SpecType SpecType `json:"spec_type"`

	// Categories defines the evaluation categories
	Categories []RubricCategory `json:"categories"`

	// PassCriteria defines what is required to pass
	PassCriteria PassCriteria `json:"pass_criteria"`
}

// RubricCategory defines an evaluation category.
type RubricCategory struct {
	// Name is the category name
	Name string `json:"name"`

	// Description explains what is evaluated
	Description string `json:"description"`

	// Weight is the category weight (0-1)
	Weight float64 `json:"weight"`

	// Criteria lists specific evaluation criteria
	Criteria []RubricCriterion `json:"criteria"`

	// Examples provide scoring examples
	Examples []ScoringExample `json:"examples,omitempty"`
}

// RubricCriterion is a specific evaluation criterion.
type RubricCriterion struct {
	// Name is the criterion name
	Name string `json:"name"`

	// Description explains what to evaluate
	Description string `json:"description"`

	// FullMarksIf describes what earns full marks
	FullMarksIf string `json:"full_marks_if"`

	// DeductIf lists conditions that reduce score
	DeductIf []string `json:"deduct_if,omitempty"`
}

// ScoringExample provides example scores for calibration.
type ScoringExample struct {
	// Score is the example score
	Score float64 `json:"score"`

	// Example describes what earns this score
	Example string `json:"example"`
}

// PassCriteria defines what is required to pass.
type PassCriteria struct {
	// MinScore is the minimum overall score to pass (0-10)
	MinScore float64 `json:"min_score"`

	// MaxCritical is the maximum critical findings allowed (-1 for unlimited)
	MaxCritical int `json:"max_critical"`

	// MaxHigh is the maximum high findings allowed (-1 for unlimited)
	MaxHigh int `json:"max_high"`

	// MaxMedium is the maximum medium findings allowed (-1 for unlimited)
	MaxMedium int `json:"max_medium"`

	// HumanReviewRequired lists what triggers human review
	HumanReviewRequired []string `json:"human_review_required,omitempty"`
}
