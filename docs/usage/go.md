# Go Library

The LEVER Go library provides types and functions for working with the framework programmatically.

## Installation

```bash
go get github.com/ProductBuildersHQ/lever
```

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/ProductBuildersHQ/lever/pkg/lever"
)

func main() {
    // Get all stages
    stages := lever.Stages()
    for _, s := range stages {
        fmt.Printf("%d. %s - %s\n", s.Order, s.Name, s.Question)
    }

    // Get complete specification
    spec := lever.NewSpec()
    data, _ := json.MarshalIndent(spec, "", "  ")
    fmt.Println(string(data))
}
```

## Core Types

### Stages

```go
// Get all stages
stages := lever.Stages()

// Get a specific stage
stage := lever.StageByID(lever.StageValue)
fmt.Println(stage.Question) // "Can I independently create outcomes?"

// All stage IDs
ids := lever.AllStageIDs()
// []StageID{"learn", "execute", "value", "enable", "replicate"}
```

### Stage Constants

```go
const (
    StageLearn     StageID = "learn"
    StageExecute   StageID = "execute"
    StageValue     StageID = "value"
    StageEnable    StageID = "enable"
    StageReplicate StageID = "replicate"
)
```

### Stage Type

```go
type Stage struct {
    ID              StageID  `json:"id"`
    Name            string   `json:"name"`
    Order           int      `json:"order"`
    Question        string   `json:"question"`
    Output          string   `json:"output"`
    Description     string   `json:"description"`
    Characteristics []string `json:"characteristics,omitempty"`
}
```

## Classical Frameworks

```go
// Get all frameworks
frameworks := lever.AllFrameworks()

// Get specific framework
bloom := lever.BloomTaxonomy()
dreyfus := lever.DreyfusModel()
shuHaRi := lever.ShuHaRi()
apprenticeship := lever.ApprenticeshipModel()
conscious := lever.ConsciousCompetenceModel()

// Get framework by ID
framework := lever.FrameworkByID(lever.FrameworkBloom)
```

### Framework Type

```go
type Framework struct {
    ID           FrameworkID      `json:"id"`
    Name         string           `json:"name"`
    Description  string           `json:"description"`
    Origin       string           `json:"origin"`
    PrimaryFocus string           `json:"primary_focus"`
    Stages       []FrameworkStage `json:"stages"`
}
```

## Mappings

```go
// Get all mappings
mappings := lever.AllMappings()

// Get specific framework mappings
bloomMappings := lever.BloomMappings()
dreyfusMappings := lever.DreyfusMappings()

// Build mapping table
table := lever.BuildMappingTable()
for _, row := range table.Rows {
    fmt.Printf("%s -> %v\n", row.LEVERStage, row.Cells)
}
```

### Mapping Types

```go
type StageMapping struct {
    LEVERStage      StageID     `json:"lever_stage"`
    FrameworkID     FrameworkID `json:"framework_id"`
    FrameworkStages []string    `json:"framework_stages"`
    Notes           string      `json:"notes,omitempty"`
}

type MappingTable struct {
    Frameworks []FrameworkID `json:"frameworks"`
    Rows       []MappingRow  `json:"rows"`
}
```

## Profiles

```go
// Reference profiles
profiles := lever.ReferenceProfiles()
productBuilder := lever.ProductBuilderProfile()
engineer := lever.EngineerProfile()

// Create custom profile
profile := lever.Profile{
    ID:          "my-domain",
    Name:        "My Domain",
    Description: "Custom LEVER interpretation",
    Domain:      "custom",
    StageInterpretations: []lever.StageInterpretation{
        {
            StageID:     lever.StageLearn,
            Title:       "Learner",
            Description: "Learning domain fundamentals",
            Evidence:    []string{"Completes training"},
        },
        // ... other stages
    },
}
```

### Profile Type

```go
type Profile struct {
    ID                   ProfileID             `json:"id"`
    Name                 string                `json:"name"`
    Description          string                `json:"description"`
    Domain               string                `json:"domain"`
    StageInterpretations []StageInterpretation `json:"stage_interpretations"`
}

type StageInterpretation struct {
    StageID      StageID      `json:"stage_id"`
    Title        string       `json:"title,omitempty"`
    Description  string       `json:"description"`
    Competencies []Competency `json:"competencies,omitempty"`
    Evidence     []string     `json:"evidence,omitempty"`
    Examples     []string     `json:"examples,omitempty"`
}
```

## Complete Specification

```go
// Get the complete spec
spec := lever.NewSpec()

// Access components
fmt.Println(spec.Version)      // "0.1.0"
fmt.Println(spec.Name)         // "LEVER Capability Framework"
fmt.Println(len(spec.Stages))  // 5
fmt.Println(len(spec.Frameworks)) // 5

// Export as JSON
data, err := json.Marshal(spec)
```

### Spec Type

```go
type Spec struct {
    Version           string             `json:"version"`
    Name              string             `json:"name"`
    Description       string             `json:"description"`
    Stages            []Stage            `json:"stages"`
    Frameworks        []Framework        `json:"frameworks"`
    Mappings          []FrameworkMappings `json:"mappings"`
    MappingTable      MappingTable       `json:"mapping_table"`
    ReferenceProfiles []Profile          `json:"reference_profiles,omitempty"`
}
```

## Stage Progression

```go
// Get progression split
prog := lever.GetStageProgression()
fmt.Println(prog.Personal)  // [learn, execute, value]
fmt.Println(prog.Leveraged) // [enable, replicate]

// Output progression
output := lever.OutputProgression[lever.StageValue]
fmt.Println(output) // "Outcomes"
```

## Embedded JSON Schema

```go
import "github.com/ProductBuildersHQ/lever/schema"

// Access embedded schemas
leverSchema := schema.LeverSchema
profileSchema := schema.ProfileSchema

// Use for validation
var spec lever.Spec
// ... validate against leverSchema
```

## Error Handling

Most functions return values directly since they work with compile-time constants:

```go
// Returns nil if not found
stage := lever.StageByID("invalid")
if stage == nil {
    // Handle unknown stage
}

framework := lever.FrameworkByID("unknown")
if framework == nil {
    // Handle unknown framework
}
```
