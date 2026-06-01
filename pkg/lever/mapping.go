package lever

// StageMapping represents a mapping from a LEVER stage to classical framework stages.
type StageMapping struct {
	// LEVERStage is the LEVER stage being mapped.
	LEVERStage StageID `json:"lever_stage"`

	// FrameworkID identifies which classical framework this maps to.
	FrameworkID FrameworkID `json:"framework_id"`

	// FrameworkStages are the corresponding stages in the classical framework.
	// Multiple stages may map to a single LEVER stage.
	FrameworkStages []string `json:"framework_stages"`

	// Notes provides context about the mapping relationship.
	Notes string `json:"notes,omitempty"`
}

// FrameworkMappings holds all mappings for a single classical framework.
type FrameworkMappings struct {
	// FrameworkID identifies the classical framework.
	FrameworkID FrameworkID `json:"framework_id"`

	// Mappings are the stage-by-stage mappings.
	Mappings []StageMapping `json:"mappings"`
}

// BloomMappings returns the LEVER → Bloom's Taxonomy mappings.
func BloomMappings() FrameworkMappings {
	return FrameworkMappings{
		FrameworkID: FrameworkBloom,
		Mappings: []StageMapping{
			{
				LEVERStage:      StageLearn,
				FrameworkID:     FrameworkBloom,
				FrameworkStages: []string{"remember", "understand"},
				Notes:           "Knowledge acquisition covers basic recall and comprehension",
			},
			{
				LEVERStage:      StageExecute,
				FrameworkID:     FrameworkBloom,
				FrameworkStages: []string{"apply"},
				Notes:           "Execution is the application of learned knowledge",
			},
			{
				LEVERStage:      StageValue,
				FrameworkID:     FrameworkBloom,
				FrameworkStages: []string{"analyze", "evaluate"},
				Notes:           "Creating value requires analysis and judgment",
			},
			{
				LEVERStage:      StageEnable,
				FrameworkID:     FrameworkBloom,
				FrameworkStages: []string{"create"},
				Notes:           "Enabling others often requires creating new approaches",
			},
			{
				LEVERStage:      StageReplicate,
				FrameworkID:     FrameworkBloom,
				FrameworkStages: []string{},
				Notes:           "Beyond Bloom - multiplication through systems is not addressed",
			},
		},
	}
}

// DreyfusMappings returns the LEVER → Dreyfus Model mappings.
func DreyfusMappings() FrameworkMappings {
	return FrameworkMappings{
		FrameworkID: FrameworkDreyfus,
		Mappings: []StageMapping{
			{
				LEVERStage:      StageLearn,
				FrameworkID:     FrameworkDreyfus,
				FrameworkStages: []string{"novice"},
				Notes:           "Novices are learning rules and fundamentals",
			},
			{
				LEVERStage:      StageExecute,
				FrameworkID:     FrameworkDreyfus,
				FrameworkStages: []string{"advanced-beginner", "competent"},
				Notes:           "Execution spans pattern recognition to planning",
			},
			{
				LEVERStage:      StageValue,
				FrameworkID:     FrameworkDreyfus,
				FrameworkStages: []string{"proficient", "expert"},
				Notes:           "Value creation requires holistic understanding and intuition",
			},
			{
				LEVERStage:      StageEnable,
				FrameworkID:     FrameworkDreyfus,
				FrameworkStages: []string{},
				Notes:           "Not directly addressed - often called 'Master' in extensions",
			},
			{
				LEVERStage:      StageReplicate,
				FrameworkID:     FrameworkDreyfus,
				FrameworkStages: []string{},
				Notes:           "Beyond Dreyfus - some extensions add 'Teacher' level",
			},
		},
	}
}

// ShuHaRiMappings returns the LEVER → Shu-Ha-Ri mappings.
func ShuHaRiMappings() FrameworkMappings {
	return FrameworkMappings{
		FrameworkID: FrameworkShuHaRi,
		Mappings: []StageMapping{
			{
				LEVERStage:      StageLearn,
				FrameworkID:     FrameworkShuHaRi,
				FrameworkStages: []string{"shu"},
				Notes:           "Shu is about learning and following the forms",
			},
			{
				LEVERStage:      StageExecute,
				FrameworkID:     FrameworkShuHaRi,
				FrameworkStages: []string{"shu"},
				Notes:           "Still within Shu - applying the learned forms",
			},
			{
				LEVERStage:      StageValue,
				FrameworkID:     FrameworkShuHaRi,
				FrameworkStages: []string{"ha"},
				Notes:           "Ha - breaking from tradition, understanding context",
			},
			{
				LEVERStage:      StageEnable,
				FrameworkID:     FrameworkShuHaRi,
				FrameworkStages: []string{"ri"},
				Notes:           "Ri practitioners often become teachers",
			},
			{
				LEVERStage:      StageReplicate,
				FrameworkID:     FrameworkShuHaRi,
				FrameworkStages: []string{"ri"},
				Notes:           "Advanced Ri - creating new schools and approaches",
			},
		},
	}
}

// ApprenticeshipMappings returns the LEVER → Apprenticeship Model mappings.
func ApprenticeshipMappings() FrameworkMappings {
	return FrameworkMappings{
		FrameworkID: FrameworkApprenticeship,
		Mappings: []StageMapping{
			{
				LEVERStage:      StageLearn,
				FrameworkID:     FrameworkApprenticeship,
				FrameworkStages: []string{"observe"},
				Notes:           "Observation is the primary learning mechanism",
			},
			{
				LEVERStage:      StageExecute,
				FrameworkID:     FrameworkApprenticeship,
				FrameworkStages: []string{"assist", "perform"},
				Notes:           "Execution spans assisted to independent work",
			},
			{
				LEVERStage:      StageValue,
				FrameworkID:     FrameworkApprenticeship,
				FrameworkStages: []string{"perform"},
				Notes:           "Full independent performance with proven results",
			},
			{
				LEVERStage:      StageEnable,
				FrameworkID:     FrameworkApprenticeship,
				FrameworkStages: []string{"lead", "teach"},
				Notes:           "Leading and teaching others",
			},
			{
				LEVERStage:      StageReplicate,
				FrameworkID:     FrameworkApprenticeship,
				FrameworkStages: []string{"teach"},
				Notes:           "Master teachers who create schools and guilds",
			},
		},
	}
}

// ConsciousCompetenceMappings returns the LEVER → Conscious Competence mappings.
func ConsciousCompetenceMappings() FrameworkMappings {
	return FrameworkMappings{
		FrameworkID: FrameworkConsciousCompetence,
		Mappings: []StageMapping{
			{
				LEVERStage:      StageLearn,
				FrameworkID:     FrameworkConsciousCompetence,
				FrameworkStages: []string{"unconscious-incompetence", "conscious-incompetence"},
				Notes:           "Learning begins with awareness of the gap",
			},
			{
				LEVERStage:      StageExecute,
				FrameworkID:     FrameworkConsciousCompetence,
				FrameworkStages: []string{"conscious-competence"},
				Notes:           "Deliberate, effortful performance",
			},
			{
				LEVERStage:      StageValue,
				FrameworkID:     FrameworkConsciousCompetence,
				FrameworkStages: []string{"unconscious-competence"},
				Notes:           "Automatic, reliable performance",
			},
			{
				LEVERStage:      StageEnable,
				FrameworkID:     FrameworkConsciousCompetence,
				FrameworkStages: []string{},
				Notes:           "Often called 'Reflective Mastery' in extensions",
			},
			{
				LEVERStage:      StageReplicate,
				FrameworkID:     FrameworkConsciousCompetence,
				FrameworkStages: []string{},
				Notes:           "Beyond the model - requires conscious articulation",
			},
		},
	}
}

// AllMappings returns all framework mappings.
func AllMappings() []FrameworkMappings {
	return []FrameworkMappings{
		BloomMappings(),
		DreyfusMappings(),
		ShuHaRiMappings(),
		ApprenticeshipMappings(),
		ConsciousCompetenceMappings(),
	}
}

// MappingTable represents the full mapping table in matrix form.
// This is useful for generating documentation and comparison tables.
type MappingTable struct {
	// Frameworks lists all frameworks in column order.
	Frameworks []FrameworkID `json:"frameworks"`

	// Rows contains one row per LEVER stage.
	Rows []MappingRow `json:"rows"`
}

// MappingRow represents one row of the mapping table.
type MappingRow struct {
	// LEVERStage is the LEVER stage for this row.
	LEVERStage StageID `json:"lever_stage"`

	// Cells maps framework ID to the corresponding stages.
	Cells map[FrameworkID][]string `json:"cells"`
}

// BuildMappingTable constructs the full mapping table.
func BuildMappingTable() MappingTable {
	allMappings := AllMappings()

	// Collect framework IDs
	frameworks := make([]FrameworkID, 0, len(allMappings))
	for _, fm := range allMappings {
		frameworks = append(frameworks, fm.FrameworkID)
	}

	// Build rows for each LEVER stage
	rows := make([]MappingRow, 0, 5)
	for _, stageID := range AllStageIDs() {
		row := MappingRow{
			LEVERStage: stageID,
			Cells:      make(map[FrameworkID][]string),
		}

		for _, fm := range allMappings {
			for _, m := range fm.Mappings {
				if m.LEVERStage == stageID {
					row.Cells[fm.FrameworkID] = m.FrameworkStages
					break
				}
			}
		}

		rows = append(rows, row)
	}

	return MappingTable{
		Frameworks: frameworks,
		Rows:       rows,
	}
}
