# LEVER Capability Framework

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/ProductBuildersHQ/lever/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/ProductBuildersHQ/lever/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/ProductBuildersHQ/lever/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/ProductBuildersHQ/lever/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/ProductBuildersHQ/lever/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/ProductBuildersHQ/lever/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/ProductBuildersHQ/lever
 [goreport-url]: https://goreportcard.com/report/github.com/ProductBuildersHQ/lever
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/ProductBuildersHQ/lever
 [docs-godoc-url]: https://pkg.go.dev/github.com/ProductBuildersHQ/lever
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://productbuildershq.com/lever
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Flever
 [loc-svg]: https://tokei.rs/b1/github/ProductBuildersHQ/lever
 [repo-url]: https://github.com/ProductBuildersHQ/lever
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/ProductBuildersHQ/lever/blob/main/LICENSE

**Learn • Execute • Value • Enable • Replicate**

LEVER is a capability multiplication framework that synthesizes classical learning models (Bloom's Taxonomy, Dreyfus Model, Shu-Ha-Ri, Apprenticeship) for the AI era.

## Core Thesis

> AI does not merely make experts more productive. It enables experts to become multipliers much sooner in their careers through agents, workflows, platforms, and standards.

## The Five Stages

| Stage | Question | Output |
|-------|----------|--------|
| **Learn** | What do I know? | Knowledge |
| **Execute** | What can I do? | Results |
| **Value** | Can I independently create outcomes? | Measurable outcomes |
| **Enable** | Can I make others successful? | Capable people |
| **Replicate** | Can I make success repeatable at scale? | Systems |

### Stage Progression

```
Personal Capability          Leveraged Capability
─────────────────────        ────────────────────
Learn → Execute → Value      Enable → Replicate
```

The first three stages develop **personal capability**. The last two stages create **leveraged capability** through people and systems.

## Framework Mappings

LEVER synthesizes and extends classical frameworks:

| LEVER | Bloom | Dreyfus | Shu-Ha-Ri | Apprenticeship |
|-------|-------|---------|-----------|----------------|
| Learn | Remember, Understand | Novice | Shu | Observe |
| Execute | Apply | Advanced Beginner, Competent | Shu | Assist, Perform |
| Value | Analyze, Evaluate | Proficient, Expert | Ha | Perform |
| Enable | Create | (Master)* | Ri | Lead, Teach |
| Replicate | — | (Teacher)* | Ri | Teach |

*Extended stages not in original model

## Installation

### Go

```bash
go get github.com/ProductBuildersHQ/lever
```

### TypeScript/JavaScript

```bash
npm install @productbuildershq/lever zod
```

## Usage

### Go Library

```go
import "github.com/ProductBuildersHQ/lever/pkg/lever"

// Get all stages
stages := lever.Stages()

// Get a specific stage
stage := lever.StageByID(lever.StageValue)

// Get the full specification
spec := lever.NewSpec()

// Build the framework mapping table
table := lever.BuildMappingTable()
```

### CLI

```bash
# Generate JSON Schema
go run ./cmd/lever schema > lever.schema.json

# Output complete spec as JSON
go run ./cmd/lever spec

# List stages
go run ./cmd/lever stages

# Output mapping table
go run ./cmd/lever mappings

# Print version
go run ./cmd/lever version
```

### TypeScript/Zod

```typescript
import {
  specSchema,
  stageSchema,
  profileSchema,
  type Spec,
  type Stage,
  type Profile,
  type StageId,
  STAGES,
  STAGE_PROGRESSION,
  OUTPUT_PROGRESSION,
} from '@productbuildershq/lever';

// Validate a spec
const result = specSchema.safeParse(data);
if (result.success) {
  const spec: Spec = result.data;
  console.log(spec.stages);
}

// Use constants
console.log(STAGES); // ['learn', 'execute', 'value', 'enable', 'replicate']
console.log(OUTPUT_PROGRESSION.learn); // 'Knowledge'

// Validate a profile
const profile = profileSchema.parse(profileData);
```

#### Pure Types (no Zod)

```typescript
import type { Spec, Stage, Profile, StageId } from '@productbuildershq/lever/types';
```

## JSON Schema

The framework provides machine-readable JSON Schema:

- [`schema/lever.schema.json`](schema/lever.schema.json) - Full specification schema
- [`schema/profile.schema.json`](schema/profile.schema.json) - Domain profile schema

### Schema Usage

```go
import "github.com/ProductBuildersHQ/lever/schema"

// Embedded schemas for validation
schemaBytes := schema.LeverSchema
profileSchemaBytes := schema.ProfileSchema
```

## Domain Profiles

LEVER can be interpreted for specific domains. Reference profiles included:

- **Product Builder** - End-to-end product ownership
- **Software Engineer** - Technical delivery

Create custom profiles by implementing `lever.Profile`:

```go
profile := lever.Profile{
    ID:          "my-domain",
    Name:        "My Domain",
    Description: "LEVER interpretation for my domain",
    Domain:      "custom",
    StageInterpretations: []lever.StageInterpretation{
        {StageID: lever.StageLearn, Description: "..."},
        // ... one per stage
    },
}
```

## Ecosystem

LEVER is the foundational theory for:

| Framework | Purpose |
|-----------|---------|
| **LEVER** | Human capability progression |
| **PBMM** | Product Builder Maturity Model |
| **ASDM** | Autonomous Software Delivery Model |
| **PRISM** | Organizational capability maturity |

```
LEVER (theory)
    ↓
PBMM (individual assessment)
ASDM (team/org assessment)
```

## Project Structure

```
lever/
├── cmd/lever/          # CLI tool
├── pkg/lever/          # Go types (source of truth)
│   ├── stage.go        # LEVER stages
│   ├── framework.go    # Classical frameworks (Bloom, Dreyfus, etc.)
│   ├── mapping.go      # Stage mappings
│   ├── profile.go      # Domain profiles
│   └── lever.go        # Complete specification
├── ts/                 # TypeScript/Zod schemas
│   ├── lever.zod.ts    # Zod schemas + types
│   ├── lever.d.ts      # Pure TypeScript types
│   └── index.ts        # Main exports
├── schema/             # Generated JSON Schema
├── dist/               # Compiled TypeScript (npm package)
└── spec.json           # Full spec as JSON
```

## Development

### Go

```bash
# Run tests
go test -v ./...

# Regenerate JSON schemas
npm run generate:schema

# Lint
golangci-lint run
```

### TypeScript

```bash
# Type-check
npm run typecheck

# Build
npm run build

# Run all tests
npm test
```

### Type Synchronization

Go structs are the source of truth. When Go types change:

1. Update Go types in `pkg/lever/`
2. Regenerate JSON Schema: `npm run generate:schema`
3. Manually update TypeScript in `ts/` to match
4. Type-check: `npm run typecheck`

## License

[License TBD]

## References

- [Bloom's Taxonomy](https://en.wikipedia.org/wiki/Bloom%27s_taxonomy) - Benjamin Bloom, 1956
- [Dreyfus Model](https://en.wikipedia.org/wiki/Dreyfus_model_of_skill_acquisition) - Stuart & Hubert Dreyfus, 1980
- [Shu-Ha-Ri](https://en.wikipedia.org/wiki/Shuhari) - Japanese martial arts tradition
- [Apprenticeship Model](https://en.wikipedia.org/wiki/Apprenticeship) - Medieval guild system
