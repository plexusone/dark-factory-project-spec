package types

// Constitution defines organization-wide standards and decisions
// that apply across all projects. Individual specs inherit from
// and may override (with justification) these defaults.
type Constitution struct {
	// Metadata about the constitution itself
	Metadata ConstitutionMetadata `json:"metadata"`

	// TechStack defines approved technologies
	TechStack TechStack `json:"tech_stack"`

	// Architecture defines architectural standards
	Architecture ArchitectureStandards `json:"architecture"`

	// FileOrganization defines project file structure standards
	FileOrganization FileOrganizationStandards `json:"file_organization"`

	// Quality defines quality and testing requirements
	Quality QualityStandards `json:"quality"`

	// Security defines security requirements
	Security SecurityStandards `json:"security"`

	// Documentation defines documentation requirements
	Documentation DocumentationStandards `json:"documentation"`

	// Process defines development process requirements
	Process ProcessStandards `json:"process"`
}

// ConstitutionMetadata contains metadata about the constitution.
type ConstitutionMetadata struct {
	// Version of this constitution (semantic versioning)
	Version string `json:"version"`

	// EffectiveDate is when this version takes effect (RFC3339)
	EffectiveDate string `json:"effective_date"`

	// ApprovedBy lists who approved this constitution
	ApprovedBy []string `json:"approved_by"`

	// LastReviewed is when this was last reviewed (RFC3339)
	LastReviewed string `json:"last_reviewed"`

	// NextReview is when this should be reviewed again (RFC3339)
	NextReview string `json:"next_review,omitempty"`
}

// TechStack defines approved technologies.
type TechStack struct {
	// Languages lists approved programming languages
	Languages []ApprovedTechnology `json:"languages"`

	// Frameworks lists approved frameworks per language
	Frameworks []ApprovedTechnology `json:"frameworks"`

	// Databases lists approved database technologies
	Databases []ApprovedTechnology `json:"databases"`

	// Infrastructure lists approved infrastructure/cloud services
	Infrastructure []ApprovedTechnology `json:"infrastructure"`

	// Libraries lists approved or required libraries
	Libraries []ApprovedTechnology `json:"libraries,omitempty"`

	// Prohibited lists explicitly banned technologies
	Prohibited []ProhibitedTechnology `json:"prohibited,omitempty"`
}

// ApprovedTechnology represents an approved technology choice.
type ApprovedTechnology struct {
	// Name of the technology
	Name string `json:"name"`

	// Category groups related technologies (e.g., "backend", "frontend", "data")
	Category string `json:"category,omitempty"`

	// Versions specifies allowed versions (e.g., ">=1.21", "^3.0.0")
	Versions string `json:"versions,omitempty"`

	// UseCase describes when to use this technology
	UseCase string `json:"use_case,omitempty"`

	// Rationale explains why this is approved
	Rationale string `json:"rationale,omitempty"`

	// Status indicates adoption status
	Status TechnologyStatus `json:"status"`

	// Owner is the team/person responsible for this choice
	Owner string `json:"owner,omitempty"`
}

// TechnologyStatus indicates the adoption status of a technology.
type TechnologyStatus string

const (
	// TechStatusPreferred is the default choice for new projects
	TechStatusPreferred TechnologyStatus = "preferred"

	// TechStatusApproved is allowed but not the default
	TechStatusApproved TechnologyStatus = "approved"

	// TechStatusLegacy is allowed for existing projects, not new
	TechStatusLegacy TechnologyStatus = "legacy"

	// TechStatusExperimental is allowed with explicit approval
	TechStatusExperimental TechnologyStatus = "experimental"

	// TechStatusDeprecated is being phased out
	TechStatusDeprecated TechnologyStatus = "deprecated"
)

// ProhibitedTechnology represents an explicitly banned technology.
type ProhibitedTechnology struct {
	// Name of the technology
	Name string `json:"name"`

	// Reason explains why it's prohibited
	Reason string `json:"reason"`

	// Alternative suggests what to use instead
	Alternative string `json:"alternative,omitempty"`

	// ExceptionProcess describes how to get an exception
	ExceptionProcess string `json:"exception_process,omitempty"`
}

// ArchitectureStandards defines architectural requirements.
type ArchitectureStandards struct {
	// Patterns lists approved architectural patterns
	Patterns []ArchitecturePattern `json:"patterns"`

	// Principles lists guiding principles
	Principles []Principle `json:"principles"`

	// APIStandards defines API design requirements
	APIStandards APIStandards `json:"api_standards,omitempty"`

	// DataStandards defines data handling requirements
	DataStandards DataStandards `json:"data_standards,omitempty"`
}

// ArchitecturePattern defines an approved pattern.
type ArchitecturePattern struct {
	// Name of the pattern
	Name string `json:"name"`

	// Description explains the pattern
	Description string `json:"description"`

	// UseCase describes when to apply this pattern
	UseCase string `json:"use_case"`

	// Reference links to documentation
	Reference string `json:"reference,omitempty"`
}

// Principle defines a guiding principle.
type Principle struct {
	// Name of the principle
	Name string `json:"name"`

	// Statement is the principle itself
	Statement string `json:"statement"`

	// Rationale explains why this matters
	Rationale string `json:"rationale,omitempty"`

	// Examples illustrate the principle
	Examples []string `json:"examples,omitempty"`
}

// APIStandards defines API design requirements.
type APIStandards struct {
	// Style is the API style (e.g., "REST", "GraphQL", "gRPC")
	Style string `json:"style"`

	// Versioning describes versioning strategy
	Versioning string `json:"versioning"`

	// Authentication describes auth requirements
	Authentication string `json:"authentication"`

	// Documentation describes API doc requirements
	Documentation string `json:"documentation"`

	// RateLimiting describes rate limit requirements
	RateLimiting string `json:"rate_limiting,omitempty"`
}

// DataStandards defines data handling requirements.
type DataStandards struct {
	// Naming describes naming conventions
	Naming string `json:"naming"`

	// Retention describes data retention requirements
	Retention string `json:"retention,omitempty"`

	// Privacy describes privacy requirements (GDPR, CCPA, etc.)
	Privacy string `json:"privacy,omitempty"`

	// Encryption describes encryption requirements
	Encryption string `json:"encryption,omitempty"`
}

// FileOrganizationStandards defines project file structure standards.
type FileOrganizationStandards struct {
	// Specs defines where machine-readable spec files live
	Specs SpecsDirectoryStandards `json:"specs"`

	// Docs defines where human-readable documentation lives
	Docs DocsDirectoryStandards `json:"docs"`

	// SpecToDocMapping defines how spec files map to doc files
	SpecToDocMapping SpecToDocMapping `json:"spec_to_doc_mapping"`

	// NamingConventions defines naming rules for files and directories
	NamingConventions NamingConventions `json:"naming_conventions"`
}

// SpecsDirectoryStandards defines machine-readable spec file locations.
// Specs are organized by feature within status directories:
//
//	specs/
//	├── roadmap.json
//	├── constitution.json
//	├── active/{feature}/
//	├── backlog/{feature}/
//	├── completed/{feature}/
//	└── archived/{feature}/
type SpecsDirectoryStandards struct {
	// RootPath is the root directory for specs (e.g., "specs/")
	RootPath string `json:"root_path"`

	// RoadmapPath is where the roadmap lives (e.g., "specs/roadmap.json")
	RoadmapPath string `json:"roadmap_path"`

	// ConstitutionPath is where the constitution lives (e.g., "specs/constitution.json")
	ConstitutionPath string `json:"constitution_path"`

	// ActivePath is the directory for active features (e.g., "specs/active/")
	ActivePath string `json:"active_path"`

	// BacklogPath is the directory for backlog features (e.g., "specs/backlog/")
	BacklogPath string `json:"backlog_path"`

	// CompletedPath is the directory for completed features (e.g., "specs/completed/")
	CompletedPath string `json:"completed_path"`

	// ArchivedPath is the directory for archived features (e.g., "specs/archived/")
	ArchivedPath string `json:"archived_path"`

	// FeatureFiles defines the standard files within each feature directory
	FeatureFiles FeatureFiles `json:"feature_files"`
}

// FeatureFiles defines the standard files within each feature directory.
type FeatureFiles struct {
	// MRDFile is the MRD filename (e.g., "mrd.json")
	MRDFile string `json:"mrd_file"`

	// PRDFile is the PRD filename (e.g., "prd.json")
	PRDFile string `json:"prd_file"`

	// EvaluationFile is the evaluation filename (e.g., "evaluation.json")
	EvaluationFile string `json:"evaluation_file"`
}

// DocsDirectoryStandards defines human-readable documentation locations.
type DocsDirectoryStandards struct {
	// RootPath is the root directory for docs (e.g., "docs/")
	RootPath string `json:"root_path"`

	// DesignPath is the subdirectory for design docs (e.g., "docs/design/")
	DesignPath string `json:"design_path"`

	// VersionedDesignPattern describes versioned design doc paths
	// e.g., "docs/design/{version}/" where version is "vX.Y.Z"
	VersionedDesignPattern string `json:"versioned_design_pattern"`

	// FeatureDesignPattern describes feature-based design doc paths
	// e.g., "docs/design/{feature}/" where feature is kebab-case
	FeatureDesignPattern string `json:"feature_design_pattern"`

	// ADRPath is the subdirectory for Architecture Decision Records
	ADRPath string `json:"adr_path,omitempty"`

	// RunbooksPath is the subdirectory for operational runbooks
	RunbooksPath string `json:"runbooks_path,omitempty"`
}

// NamingConventions defines naming rules for files and directories.
type NamingConventions struct {
	// VersionFormat is the regex for version directories (e.g., "^v\\d+\\.\\d+\\.\\d+$")
	VersionFormat string `json:"version_format"`

	// VersionExample shows a valid version (e.g., "v1.2.3")
	VersionExample string `json:"version_example"`

	// FeatureFormat is the regex for feature directories (e.g., "^[a-z][a-z0-9]*(-[a-z0-9]+)*$")
	FeatureFormat string `json:"feature_format"`

	// FeatureExample shows a valid feature name (e.g., "user-authentication")
	FeatureExample string `json:"feature_example"`

	// SpecFileFormat is the naming pattern for spec files (e.g., "{type}-{id}.json")
	SpecFileFormat string `json:"spec_file_format"`

	// SpecFileExample shows a valid spec filename (e.g., "mrd-2025-001.json")
	SpecFileExample string `json:"spec_file_example"`
}

// SpecToDocMapping defines how machine spec files map to human doc files.
type SpecToDocMapping struct {
	// Strategy determines how specs are organized in docs
	// "feature" = docs/design/{feature}/{spec-id}.md
	// "version" = docs/design/{version}/{spec-id}.md
	// "mirror"  = docs/specs/{type}/{spec-id}.md (mirrors spec structure)
	Strategy DocOrganizationStrategy `json:"strategy"`

	// FeatureField is the JSON path in the spec that contains the feature name
	// e.g., "metadata.feature" - only used when strategy is "feature"
	FeatureField string `json:"feature_field,omitempty"`

	// VersionField is the JSON path in the spec that contains the version
	// e.g., "overview.target_release" - only used when strategy is "version"
	VersionField string `json:"version_field,omitempty"`

	// DocFileExtension is the extension for doc files (e.g., ".md")
	DocFileExtension string `json:"doc_file_extension"`

	// IncludeSpecTypePrefix includes the spec type in doc filename
	// true: mrd-2025-001.md, false: 2025-001.md
	IncludeSpecTypePrefix bool `json:"include_spec_type_prefix"`

	// GenerateIndex creates an index.md in each directory
	GenerateIndex bool `json:"generate_index"`
}

// DocOrganizationStrategy determines how docs are organized.
type DocOrganizationStrategy string

const (
	// DocStrategyFeature organizes docs by feature: docs/design/{feature}/{spec}.md
	DocStrategyFeature DocOrganizationStrategy = "feature"

	// DocStrategyVersion organizes docs by version: docs/design/{version}/{spec}.md
	DocStrategyVersion DocOrganizationStrategy = "version"

	// DocStrategyMirror mirrors spec structure: docs/specs/{type}/{spec}.md
	DocStrategyMirror DocOrganizationStrategy = "mirror"
)

// QualityStandards defines quality requirements.
type QualityStandards struct {
	// Testing defines testing requirements
	Testing TestingStandards `json:"testing"`

	// CodeReview defines review requirements
	CodeReview CodeReviewStandards `json:"code_review"`

	// Metrics defines quality metrics
	Metrics []QualityMetric `json:"metrics,omitempty"`
}

// TestingStandards defines testing requirements.
type TestingStandards struct {
	// UnitTestCoverage is minimum unit test coverage (e.g., "80%")
	UnitTestCoverage string `json:"unit_test_coverage"`

	// IntegrationTestRequired indicates if integration tests are required
	IntegrationTestRequired bool `json:"integration_test_required"`

	// E2ETestRequired indicates if E2E tests are required
	E2ETestRequired bool `json:"e2e_test_required"`

	// PerformanceTestRequired indicates if perf tests are required
	PerformanceTestRequired bool `json:"performance_test_required"`

	// SecurityTestRequired indicates if security tests are required
	SecurityTestRequired bool `json:"security_test_required"`

	// TestFrameworks lists approved testing frameworks
	TestFrameworks []string `json:"test_frameworks,omitempty"`
}

// CodeReviewStandards defines review requirements.
type CodeReviewStandards struct {
	// RequiredApprovals is minimum approvals needed
	RequiredApprovals int `json:"required_approvals"`

	// RequireOwnerApproval requires codeowner approval
	RequireOwnerApproval bool `json:"require_owner_approval"`

	// AutomatedChecks lists required automated checks
	AutomatedChecks []string `json:"automated_checks"`

	// MaxPRSize is max lines changed per PR (0 = no limit)
	MaxPRSize int `json:"max_pr_size,omitempty"`
}

// QualityMetric defines a quality metric.
type QualityMetric struct {
	// Name of the metric
	Name string `json:"name"`

	// Description explains the metric
	Description string `json:"description"`

	// Target is the target value
	Target string `json:"target"`

	// Measurement explains how to measure
	Measurement string `json:"measurement"`
}

// SecurityStandards defines security requirements.
type SecurityStandards struct {
	// Authentication defines auth requirements
	Authentication AuthenticationStandards `json:"authentication"`

	// Authorization defines authz requirements
	Authorization AuthorizationStandards `json:"authorization"`

	// DataProtection defines data protection requirements
	DataProtection DataProtectionStandards `json:"data_protection"`

	// Compliance lists compliance requirements
	Compliance []string `json:"compliance,omitempty"`

	// VulnerabilityScanning defines scanning requirements
	VulnerabilityScanning VulnerabilityScanningStandards `json:"vulnerability_scanning"`
}

// AuthenticationStandards defines authentication requirements.
type AuthenticationStandards struct {
	// Methods lists approved auth methods
	Methods []string `json:"methods"`

	// MFARequired indicates if MFA is required
	MFARequired bool `json:"mfa_required"`

	// SessionTimeout is max session duration
	SessionTimeout string `json:"session_timeout,omitempty"`

	// PasswordPolicy describes password requirements
	PasswordPolicy string `json:"password_policy,omitempty"`
}

// AuthorizationStandards defines authorization requirements.
type AuthorizationStandards struct {
	// Model is the authz model (e.g., "RBAC", "ABAC", "Cedar")
	Model string `json:"model"`

	// DefaultDeny indicates if default policy is deny
	DefaultDeny bool `json:"default_deny"`

	// AuditRequired indicates if access must be audited
	AuditRequired bool `json:"audit_required"`
}

// DataProtectionStandards defines data protection requirements.
type DataProtectionStandards struct {
	// EncryptionAtRest requires encryption at rest
	EncryptionAtRest bool `json:"encryption_at_rest"`

	// EncryptionInTransit requires encryption in transit
	EncryptionInTransit bool `json:"encryption_in_transit"`

	// PIIHandling describes PII handling requirements
	PIIHandling string `json:"pii_handling,omitempty"`

	// SecretManagement describes secret management requirements
	SecretManagement string `json:"secret_management"`
}

// VulnerabilityScanningStandards defines scanning requirements.
type VulnerabilityScanningStandards struct {
	// DependencyScanning requires dependency vulnerability scanning
	DependencyScanning bool `json:"dependency_scanning"`

	// ContainerScanning requires container image scanning
	ContainerScanning bool `json:"container_scanning"`

	// SASTRequired requires static analysis
	SASTRequired bool `json:"sast_required"`

	// DASTRequired requires dynamic analysis
	DASTRequired bool `json:"dast_required"`

	// MaxCriticalVulnerabilities is max allowed critical vulns (0 = none)
	MaxCriticalVulnerabilities int `json:"max_critical_vulnerabilities"`
}

// DocumentationStandards defines documentation requirements.
type DocumentationStandards struct {
	// READMERequired requires README for all projects
	READMERequired bool `json:"readme_required"`

	// APIDocRequired requires API documentation
	APIDocRequired bool `json:"api_doc_required"`

	// ArchitectureDocRequired requires architecture docs
	ArchitectureDocRequired bool `json:"architecture_doc_required"`

	// RunbookRequired requires operational runbooks
	RunbookRequired bool `json:"runbook_required"`

	// ChangelogRequired requires changelog maintenance
	ChangelogRequired bool `json:"changelog_required"`

	// Format describes documentation format requirements
	Format string `json:"format,omitempty"`
}

// ProcessStandards defines development process requirements.
type ProcessStandards struct {
	// BranchingStrategy describes git branching strategy
	BranchingStrategy string `json:"branching_strategy"`

	// CommitConvention describes commit message format
	CommitConvention string `json:"commit_convention"`

	// ReleaseProcess describes release process
	ReleaseProcess string `json:"release_process"`

	// IncidentResponse describes incident response process
	IncidentResponse string `json:"incident_response,omitempty"`

	// OnCallRequired indicates if on-call is required
	OnCallRequired bool `json:"on_call_required"`
}

// ConstitutionOverride allows a spec to override constitution defaults.
type ConstitutionOverride struct {
	// ConstitutionVersion is the version being overridden
	ConstitutionVersion string `json:"constitution_version"`

	// Overrides lists specific overrides
	Overrides []Override `json:"overrides"`
}

// Override represents a specific deviation from the constitution.
type Override struct {
	// Path is the JSON path being overridden (e.g., "tech_stack.languages")
	Path string `json:"path"`

	// OriginalValue is what the constitution specifies
	OriginalValue string `json:"original_value"`

	// OverrideValue is what this spec uses instead
	OverrideValue string `json:"override_value"`

	// Justification explains why the override is needed
	Justification string `json:"justification"`

	// ApprovedBy lists who approved this override
	ApprovedBy []string `json:"approved_by"`

	// ApprovedAt is when the override was approved (RFC3339)
	ApprovedAt string `json:"approved_at"`

	// ExpiresAt is when this override expires (RFC3339, optional)
	ExpiresAt string `json:"expires_at,omitempty"`
}
