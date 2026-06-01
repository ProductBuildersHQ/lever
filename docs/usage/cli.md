# CLI

The LEVER CLI provides command-line access to the framework specification.

## Installation

```bash
go install github.com/ProductBuildersHQ/lever/cmd/lever@latest
```

Or build from source:

```bash
git clone https://github.com/ProductBuildersHQ/lever
cd lever
go build -o lever ./cmd/lever
```

## Commands

### lever stages

List all LEVER stages with their questions and outputs:

```bash
$ lever stages
1. Learn (learn)
   Question: Can I follow instructions?
   Output:   Knowledge

2. Execute (execute)
   Question: Can I complete tasks?
   Output:   Work

3. Value (value)
   Question: Can I independently create outcomes?
   Output:   Outcomes

4. Enable (enable)
   Question: Can I develop others?
   Output:   People

5. Replicate (replicate)
   Question: Can I multiply my impact?
   Output:   Systems
```

### lever spec

Output the complete LEVER specification as JSON:

```bash
$ lever spec > lever-spec.json

$ lever spec | jq '.stages[2]'
{
  "id": "value",
  "name": "Value",
  "order": 3,
  "question": "Can I independently create outcomes?",
  "output": "Outcomes",
  "description": "Independently creating outcomes that matter..."
}
```

### lever schema

Generate JSON Schema for the LEVER specification:

```bash
$ lever schema > lever.schema.json

$ lever schema | jq '.properties.stages'
{
  "items": { "$ref": "#/$defs/Stage" },
  "type": "array"
}
```

### lever profile-schema

Generate JSON Schema for LEVER profiles:

```bash
$ lever profile-schema > profile.schema.json
```

### lever mappings

Output the framework mapping table as JSON:

```bash
$ lever mappings | jq '.rows[0]'
{
  "lever_stage": "learn",
  "cells": {
    "bloom": ["remember", "understand"],
    "dreyfus": ["novice"],
    "shu-ha-ri": ["shu"],
    "apprenticeship": ["observe"],
    "conscious-competence": ["unconscious-incompetence", "conscious-incompetence"]
  }
}
```

### lever version

Print the specification version:

```bash
$ lever version
0.1.0
```

## Usage Examples

### Extract Stage Questions

```bash
$ lever spec | jq -r '.stages[] | "\(.order). \(.question)"'
1. Can I follow instructions?
2. Can I complete tasks?
3. Can I independently create outcomes?
4. Can I develop others?
5. Can I multiply my impact?
```

### Get Framework Mapping for a Stage

```bash
$ lever mappings | jq '.rows[] | select(.lever_stage == "value")'
{
  "lever_stage": "value",
  "cells": {
    "bloom": ["analyze", "evaluate"],
    "dreyfus": ["proficient", "expert"],
    "shu-ha-ri": ["ha"],
    "apprenticeship": ["perform"],
    "conscious-competence": ["unconscious-competence"]
  }
}
```

### List All Classical Frameworks

```bash
$ lever spec | jq -r '.frameworks[] | "\(.name) - \(.primary_focus)"'
Bloom's Taxonomy - Cognitive complexity levels
Dreyfus Model - Skill acquisition stages
Shu-Ha-Ri - Japanese mastery progression
Apprenticeship Model - Traditional craft learning
Conscious Competence Model - Awareness of ability
```

### Export Reference Profiles

```bash
$ lever spec | jq '.reference_profiles[] | {id, name, domain}'
{
  "id": "product-builder",
  "name": "Product Builder",
  "domain": "product"
}
{
  "id": "engineer",
  "name": "Software Engineer",
  "domain": "engineering"
}
```

### Generate Markdown Table

```bash
$ lever spec | jq -r '
  ["Stage", "Question", "Output"],
  ["---", "---", "---"],
  (.stages[] | [.name, .question, .output])
  | @tsv' | column -t -s $'\t'

Stage      Question                              Output
---        ---                                   ---
Learn      Can I follow instructions?            Knowledge
Execute    Can I complete tasks?                 Work
Value      Can I independently create outcomes?  Outcomes
Enable     Can I develop others?                 People
Replicate  Can I multiply my impact?             Systems
```

## Integration

### CI/CD Pipeline

```yaml
# .github/workflows/validate.yml
jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - run: go install github.com/ProductBuildersHQ/lever/cmd/lever@latest
      - run: lever spec > spec.json
      - run: lever schema > schema.json
```

### Validation Script

```bash
#!/bin/bash
# validate-profile.sh

PROFILE_FILE=$1
SCHEMA_FILE=$(mktemp)

lever profile-schema > "$SCHEMA_FILE"

# Validate with ajv or similar
npx ajv validate -s "$SCHEMA_FILE" -d "$PROFILE_FILE"

rm "$SCHEMA_FILE"
```

## Help

```bash
$ lever
lever - LEVER Capability Framework CLI

Usage:
  lever <command>

Commands:
  schema          Generate JSON Schema for lever.Spec (stdout)
  profile-schema  Generate JSON Schema for lever.Profile (stdout)
  spec            Output the complete LEVER spec as JSON (stdout)
  stages          List LEVER stages
  mappings        Output the mapping table as JSON
  version         Print the spec version
```
