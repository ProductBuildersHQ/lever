# CLAUDE.md - LEVER Capability Framework

Project-specific instructions for Claude Code.

## Project Overview

LEVER is a capability multiplication framework that synthesizes classical learning models (Bloom, Dreyfus, Shu-Ha-Ri, Apprenticeship) for the AI era.

**LEVER = Learn, Execute, Value, Enable, Replicate**

## Architecture

This is a **polyglot project** with Go as the source of truth:

```
Go Structs (source of truth)
    ↓
JSON Schema (generated)
    ↓
TypeScript/Zod (manually maintained, must match Go)
```

### Directory Structure

```
lever/
├── pkg/lever/           # Go types (SOURCE OF TRUTH)
│   ├── stage.go         # LEVER stages
│   ├── framework.go     # Classical frameworks (Bloom, Dreyfus, etc.)
│   ├── mapping.go       # Stage mappings
│   ├── profile.go       # Domain profiles
│   └── lever.go         # Complete specification
├── cmd/lever/           # CLI for schema generation
├── schema/              # Generated JSON Schema
│   ├── lever.schema.json
│   └── profile.schema.json
├── ts/                  # TypeScript/Zod (manually maintained)
│   ├── lever.zod.ts     # Zod schemas + types
│   ├── lever.d.ts       # Pure TypeScript types
│   ├── profile.zod.ts   # Profile exports
│   └── index.ts         # Main exports
├── dist/                # Compiled TypeScript (gitignored)
└── docs/                # Documentation
    └── IDEATION_CHAT.md # Original ideation conversation
```

## Development Workflow

### When Modifying Go Types

1. Edit Go structs in `pkg/lever/`
2. Run tests: `go test -v ./...`
3. Regenerate JSON Schema: `npm run generate:schema`
4. **Manually update** TypeScript in `ts/` to match Go changes
5. Type-check TypeScript: `npm run typecheck`
6. Build: `npm run build`

### Commands

```bash
# Go
go test -v ./...              # Run Go tests
go run ./cmd/lever schema     # Output JSON Schema to stdout
go run ./cmd/lever spec       # Output complete spec as JSON
go run ./cmd/lever stages     # List LEVER stages
golangci-lint run             # Lint Go code

# TypeScript
npm run typecheck             # Type-check without emitting
npm run build                 # Compile TypeScript to dist/
npm run generate:schema       # Regenerate JSON schemas from Go
npm test                      # Run all tests (Go + TypeScript)
```

## Type Synchronization

**CRITICAL:** When Go types change, TypeScript must be manually updated.

### Go → TypeScript Mapping

| Go Type | TypeScript Type | Zod Schema |
|---------|-----------------|------------|
| `StageID` | `StageId` | `stageIdSchema` |
| `Stage` | `Stage` | `stageSchema` |
| `FrameworkID` | `FrameworkId` | `frameworkIdSchema` |
| `Framework` | `Framework` | `frameworkSchema` |
| `StageMapping` | `StageMapping` | `stageMappingSchema` |
| `FrameworkMappings` | `FrameworkMappings` | `frameworkMappingsSchema` |
| `Profile` | `Profile` | `profileSchema` |
| `Spec` | `Spec` | `specSchema` |

### JSON Tag → TypeScript Property

Go uses `json:"snake_case"` tags. TypeScript properties match:

```go
// Go
type Stage struct {
    ID          StageID  `json:"id"`
    PrimaryFocus string  `json:"primary_focus"`
}
```

```typescript
// TypeScript
interface Stage {
  id: StageId;
  primary_focus: string;
}
```

## JSON Schema Generation

JSON Schema is generated from Go types using `invopop/jsonschema`:

```bash
# Generate lever.schema.json
go run ./cmd/lever schema > schema/lever.schema.json

# Generate profile.schema.json
go run ./cmd/lever profile-schema > schema/profile.schema.json
```

The schemas use `$defs` for type definitions and `$ref` for references.

## TypeScript/Zod Design

### File Organization

- `lever.zod.ts` - All Zod schemas with inferred types
- `lever.d.ts` - Pure TypeScript interfaces (for users who don't want Zod)
- `profile.zod.ts` - Re-exports profile-related schemas
- `index.ts` - Main entry point

### Exports

```typescript
// Import Zod schemas
import { specSchema, stageSchema, profileSchema } from '@productbuildershq/lever';

// Import types
import type { Spec, Stage, Profile, StageId } from '@productbuildershq/lever';

// Import constants
import { STAGES, STAGE_PROGRESSION, OUTPUT_PROGRESSION } from '@productbuildershq/lever';
```

## Adding New Types

1. **Add Go type** in `pkg/lever/`
2. **Add to Spec** if it should be in the main specification
3. **Write tests** in `*_test.go`
4. **Regenerate schema**: `npm run generate:schema`
5. **Add TypeScript**:
   - Add Zod schema in `ts/lever.zod.ts`
   - Add interface in `ts/lever.d.ts`
   - Export from `ts/index.ts` if needed
6. **Type-check**: `npm run typecheck`

## Testing

```bash
# All tests
npm test

# Go only
go test -v ./...

# TypeScript type-check only
npm run typecheck

# Lint Go
golangci-lint run
```

## Publishing

The npm package includes:

- `dist/` - Compiled TypeScript
- `schema/` - JSON Schema files

```bash
npm run build        # Compile TypeScript
npm publish          # Publish to npm
```

## Related Projects

LEVER is the foundational theory for:

| Project | Purpose |
|---------|---------|
| **PBMM** | Product Builder Maturity Model |
| **ASDM** | Autonomous Software Delivery Model |
| **PRISM** | Organizational capability maturity |

## Core Concepts

### The Five Stages

| Stage | Question | Output |
|-------|----------|--------|
| Learn | What do I know? | Knowledge |
| Execute | What can I do? | Results |
| Value | Can I independently create outcomes? | Outcomes |
| Enable | Can I make others successful? | People |
| Replicate | Can I make success repeatable at scale? | Systems |

### Stage Progression

```
Personal Capability          Leveraged Capability
─────────────────────        ────────────────────
Learn → Execute → Value      Enable → Replicate
```

### Classical Framework Mappings

LEVER maps to Bloom, Dreyfus, Shu-Ha-Ri, Apprenticeship, and Conscious Competence models. See `pkg/lever/mapping.go` for the formal mappings.
