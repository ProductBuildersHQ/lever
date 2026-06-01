package lever

// Version is the current version of the LEVER specification.
const Version = "0.1.0"

// Spec represents the complete LEVER Capability Framework specification.
// This is the root type for JSON Schema generation.
type Spec struct {
	// Version is the spec version (semver).
	Version string `json:"version"`

	// Name is the framework name.
	Name string `json:"name"`

	// Description provides an overview of the framework.
	Description string `json:"description"`

	// Stages are the five LEVER stages.
	Stages []Stage `json:"stages"`

	// Frameworks are the classical learning/mastery frameworks.
	Frameworks []Framework `json:"frameworks"`

	// Mappings connect LEVER stages to classical framework stages.
	Mappings []FrameworkMappings `json:"mappings"`

	// MappingTable provides the matrix view of all mappings.
	MappingTable MappingTable `json:"mapping_table"`

	// ReferenceProfiles demonstrate domain-specific interpretations.
	ReferenceProfiles []Profile `json:"reference_profiles,omitempty"`
}

// NewSpec creates a new LEVER specification with all canonical data.
func NewSpec() Spec {
	return Spec{
		Version: Version,
		Name:    "LEVER Capability Framework",
		Description: "LEVER (Learn, Execute, Value, Enable, Replicate) is a capability " +
			"multiplication framework that synthesizes classical learning models " +
			"(Bloom's Taxonomy, Dreyfus Model, Shu-Ha-Ri, Apprenticeship) for the AI era. " +
			"It describes how individuals progress from acquiring knowledge to creating value, " +
			"enabling others, and ultimately replicating capability through systems, automation, " +
			"AI agents, platforms, and institutions.",
		Stages:            Stages(),
		Frameworks:        AllFrameworks(),
		Mappings:          AllMappings(),
		MappingTable:      BuildMappingTable(),
		ReferenceProfiles: ReferenceProfiles(),
	}
}

// CoreThesis returns the central thesis of LEVER.
func CoreThesis() string {
	return "AI does not merely make experts more productive. " +
		"It enables experts to become multipliers much sooner in their careers " +
		"through agents, workflows, platforms, and standards."
}

// StageProgression describes the output evolution through LEVER stages.
type StageProgression struct {
	// Personal represents the first three stages (Learn, Execute, Value).
	// Output: Personal capability
	Personal []StageID `json:"personal"`

	// Leveraged represents the last two stages (Enable, Replicate).
	// Output: Multiplied capability through others and systems
	Leveraged []StageID `json:"leveraged"`
}

// GetStageProgression returns the two-part stage progression.
func GetStageProgression() StageProgression {
	return StageProgression{
		Personal:  []StageID{StageLearn, StageExecute, StageValue},
		Leveraged: []StageID{StageEnable, StageReplicate},
	}
}

// OutputProgression describes what each stage produces.
var OutputProgression = map[StageID]string{
	StageLearn:     "Knowledge",
	StageExecute:   "Work",
	StageValue:     "Outcomes",
	StageEnable:    "People",
	StageReplicate: "Systems",
}
