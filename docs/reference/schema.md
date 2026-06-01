# JSON Schema Reference

LEVER provides JSON Schemas for validation and tooling integration.

## Available Schemas

| Schema | Description | URL |
|--------|-------------|-----|
| `lever.schema.json` | Complete LEVER specification | `schema/lever.schema.json` |
| `profile.schema.json` | Domain profile definition | `schema/profile.schema.json` |

## Generating Schemas

Schemas are generated from Go types:

```bash
# Generate spec schema
lever schema > schema/lever.schema.json

# Generate profile schema
lever profile-schema > schema/profile.schema.json
```

## Schema Structure

### lever.schema.json

The main specification schema defines:

```json
{
  "$id": "https://productbuildershq.com/schemas/lever.schema.json",
  "title": "LEVER Capability Framework",
  "type": "object",
  "properties": {
    "version": { "type": "string" },
    "name": { "type": "string" },
    "description": { "type": "string" },
    "stages": { "type": "array", "items": { "$ref": "#/$defs/Stage" } },
    "frameworks": { "type": "array", "items": { "$ref": "#/$defs/Framework" } },
    "mappings": { "type": "array", "items": { "$ref": "#/$defs/FrameworkMappings" } },
    "mapping_table": { "$ref": "#/$defs/MappingTable" },
    "reference_profiles": { "type": "array", "items": { "$ref": "#/$defs/Profile" } }
  }
}
```

### Key Definitions

#### Stage

```json
{
  "$defs": {
    "Stage": {
      "type": "object",
      "properties": {
        "id": { "$ref": "#/$defs/StageID" },
        "name": { "type": "string" },
        "order": { "type": "integer" },
        "question": { "type": "string" },
        "output": { "type": "string" },
        "description": { "type": "string" },
        "characteristics": { "type": "array", "items": { "type": "string" } }
      },
      "required": ["id", "name", "order", "question", "output", "description"]
    }
  }
}
```

#### StageID

```json
{
  "$defs": {
    "StageID": {
      "type": "string",
      "enum": ["learn", "execute", "value", "enable", "replicate"]
    }
  }
}
```

#### FrameworkID

```json
{
  "$defs": {
    "FrameworkID": {
      "type": "string",
      "enum": ["bloom", "dreyfus", "shu-ha-ri", "apprenticeship", "conscious-competence"]
    }
  }
}
```

### profile.schema.json

The profile schema for domain-specific interpretations:

```json
{
  "$id": "https://productbuildershq.com/schemas/profile.schema.json",
  "title": "LEVER Profile",
  "type": "object",
  "properties": {
    "id": { "type": "string" },
    "name": { "type": "string" },
    "description": { "type": "string" },
    "domain": { "type": "string" },
    "stage_interpretations": {
      "type": "array",
      "items": { "$ref": "#/$defs/StageInterpretation" }
    }
  },
  "required": ["id", "name", "description", "domain", "stage_interpretations"]
}
```

## Validation

### Using ajv (Node.js)

```bash
npm install -g ajv-cli

# Validate a spec
ajv validate -s schema/lever.schema.json -d my-spec.json

# Validate a profile
ajv validate -s schema/profile.schema.json -d my-profile.json
```

### Using Go

```go
import (
    "github.com/ProductBuildersHQ/lever/schema"
    "github.com/xeipuuv/gojsonschema"
)

func validateProfile(profileJSON []byte) error {
    schemaLoader := gojsonschema.NewStringLoader(string(schema.ProfileSchema))
    docLoader := gojsonschema.NewBytesLoader(profileJSON)

    result, err := gojsonschema.Validate(schemaLoader, docLoader)
    if err != nil {
        return err
    }

    if !result.Valid() {
        for _, err := range result.Errors() {
            fmt.Printf("- %s\n", err)
        }
        return errors.New("validation failed")
    }

    return nil
}
```

### Using TypeScript/Zod

```typescript
import { profileSchema } from './lever.zod';

const result = profileSchema.safeParse(profileData);
if (!result.success) {
  console.error(result.error.format());
}
```

## Embedding in Go

The schema package provides embedded schemas:

```go
import "github.com/ProductBuildersHQ/lever/schema"

// Access embedded schema files
var leverSchema = schema.LeverSchema
var profileSchema = schema.ProfileSchema
```

## IDE Integration

### VS Code

Add to `.vscode/settings.json`:

```json
{
  "json.schemas": [
    {
      "fileMatch": ["**/lever-spec.json"],
      "url": "./schema/lever.schema.json"
    },
    {
      "fileMatch": ["**/profiles/*.json", "**/*-profile.json"],
      "url": "./schema/profile.schema.json"
    }
  ]
}
```

### JetBrains IDEs

1. Open Settings → Languages & Frameworks → Schemas and DTDs → JSON Schema Mappings
2. Add mappings for `lever.schema.json` and `profile.schema.json`

## Schema URLs

Production schema URLs:

- `https://productbuildershq.com/schemas/lever.schema.json`
- `https://productbuildershq.com/schemas/profile.schema.json`

Use these in JSON files for validation:

```json
{
  "$schema": "https://productbuildershq.com/schemas/profile.schema.json",
  "id": "my-profile",
  "name": "My Profile",
  ...
}
```
