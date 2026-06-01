package lever

import (
	"encoding/json"
	"testing"
)

func TestStages(t *testing.T) {
	stages := Stages()
	if len(stages) != 5 {
		t.Errorf("expected 5 stages, got %d", len(stages))
	}

	// Verify order
	expectedOrder := []StageID{StageLearn, StageExecute, StageValue, StageEnable, StageReplicate}
	for i, s := range stages {
		if s.ID != expectedOrder[i] {
			t.Errorf("stage %d: expected %s, got %s", i, expectedOrder[i], s.ID)
		}
		if s.Order != i+1 {
			t.Errorf("stage %s: expected order %d, got %d", s.ID, i+1, s.Order)
		}
	}
}

func TestStageByID(t *testing.T) {
	tests := []struct {
		id       StageID
		wantName string
		wantNil  bool
	}{
		{StageLearn, "Learn", false},
		{StageExecute, "Execute", false},
		{StageValue, "Value", false},
		{StageEnable, "Enable", false},
		{StageReplicate, "Replicate", false},
		{"invalid", "", true},
	}

	for _, tt := range tests {
		s := StageByID(tt.id)
		if tt.wantNil {
			if s != nil {
				t.Errorf("StageByID(%q): expected nil, got %v", tt.id, s)
			}
		} else {
			if s == nil {
				t.Errorf("StageByID(%q): unexpected nil", tt.id)
			} else if s.Name != tt.wantName {
				t.Errorf("StageByID(%q): expected name %q, got %q", tt.id, tt.wantName, s.Name)
			}
		}
	}
}

func TestAllFrameworks(t *testing.T) {
	frameworks := AllFrameworks()
	if len(frameworks) != 5 {
		t.Errorf("expected 5 frameworks, got %d", len(frameworks))
	}

	expectedIDs := []FrameworkID{
		FrameworkBloom,
		FrameworkDreyfus,
		FrameworkShuHaRi,
		FrameworkApprenticeship,
		FrameworkConsciousCompetence,
	}

	for i, f := range frameworks {
		if f.ID != expectedIDs[i] {
			t.Errorf("framework %d: expected %s, got %s", i, expectedIDs[i], f.ID)
		}
		if len(f.Stages) == 0 {
			t.Errorf("framework %s: no stages defined", f.ID)
		}
	}
}

func TestAllMappings(t *testing.T) {
	mappings := AllMappings()
	if len(mappings) != 5 {
		t.Errorf("expected 5 framework mappings, got %d", len(mappings))
	}

	// Each framework mapping should have 5 stage mappings (one per LEVER stage)
	for _, fm := range mappings {
		if len(fm.Mappings) != 5 {
			t.Errorf("framework %s: expected 5 stage mappings, got %d", fm.FrameworkID, len(fm.Mappings))
		}
	}
}

func TestBuildMappingTable(t *testing.T) {
	table := BuildMappingTable()

	if len(table.Frameworks) != 5 {
		t.Errorf("expected 5 frameworks in table, got %d", len(table.Frameworks))
	}

	if len(table.Rows) != 5 {
		t.Errorf("expected 5 rows (one per LEVER stage), got %d", len(table.Rows))
	}

	// Each row should have cells for each framework
	for _, row := range table.Rows {
		if len(row.Cells) != 5 {
			t.Errorf("row %s: expected 5 cells, got %d", row.LEVERStage, len(row.Cells))
		}
	}
}

func TestNewSpec(t *testing.T) {
	spec := NewSpec()

	if spec.Version != Version {
		t.Errorf("expected version %s, got %s", Version, spec.Version)
	}

	if len(spec.Stages) != 5 {
		t.Errorf("expected 5 stages, got %d", len(spec.Stages))
	}

	if len(spec.Frameworks) != 5 {
		t.Errorf("expected 5 frameworks, got %d", len(spec.Frameworks))
	}

	if len(spec.Mappings) != 5 {
		t.Errorf("expected 5 mappings, got %d", len(spec.Mappings))
	}
}

func TestSpecJSON(t *testing.T) {
	spec := NewSpec()

	data, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("failed to marshal spec: %v", err)
	}

	var decoded Spec
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal spec: %v", err)
	}

	if decoded.Version != spec.Version {
		t.Errorf("round-trip: version mismatch: %s != %s", decoded.Version, spec.Version)
	}

	if len(decoded.Stages) != len(spec.Stages) {
		t.Errorf("round-trip: stage count mismatch: %d != %d", len(decoded.Stages), len(spec.Stages))
	}
}

func TestReferenceProfiles(t *testing.T) {
	profiles := ReferenceProfiles()
	if len(profiles) < 2 {
		t.Errorf("expected at least 2 reference profiles, got %d", len(profiles))
	}

	// Each profile should have interpretations for all 5 stages
	for _, p := range profiles {
		if len(p.StageInterpretations) != 5 {
			t.Errorf("profile %s: expected 5 stage interpretations, got %d", p.ID, len(p.StageInterpretations))
		}
	}
}

func TestStageProgression(t *testing.T) {
	prog := GetStageProgression()

	if len(prog.Personal) != 3 {
		t.Errorf("expected 3 personal stages, got %d", len(prog.Personal))
	}

	if len(prog.Leveraged) != 2 {
		t.Errorf("expected 2 leveraged stages, got %d", len(prog.Leveraged))
	}

	// Verify personal stages
	expectedPersonal := []StageID{StageLearn, StageExecute, StageValue}
	for i, s := range prog.Personal {
		if s != expectedPersonal[i] {
			t.Errorf("personal stage %d: expected %s, got %s", i, expectedPersonal[i], s)
		}
	}

	// Verify leveraged stages
	expectedLeveraged := []StageID{StageEnable, StageReplicate}
	for i, s := range prog.Leveraged {
		if s != expectedLeveraged[i] {
			t.Errorf("leveraged stage %d: expected %s, got %s", i, expectedLeveraged[i], s)
		}
	}
}
