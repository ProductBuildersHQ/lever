# TypeScript / Zod

LEVER provides TypeScript types and Zod schemas for runtime validation.

## Installation

```bash
npm install zod
```

Copy the LEVER type files to your project:

- `ts/lever.zod.ts` — Zod schemas with inferred types
- `ts/lever.d.ts` — Pure TypeScript interfaces

## Using Zod Schemas

```typescript
import {
  stageIdSchema,
  specSchema,
  profileSchema,
  STAGES,
  STAGE_PROGRESSION
} from './lever.zod';

// Validate a stage ID
const result = stageIdSchema.safeParse('value');
if (result.success) {
  console.log(`Valid stage: ${result.data}`);
}

// Validate a complete spec
const specData = await fetch('/api/lever-spec').then(r => r.json());
const spec = specSchema.parse(specData);
console.log(`LEVER v${spec.version}`);

// Use constants
console.log(STAGES); // ['learn', 'execute', 'value', 'enable', 'replicate']
console.log(STAGE_PROGRESSION.personal);  // ['learn', 'execute', 'value']
console.log(STAGE_PROGRESSION.leveraged); // ['enable', 'replicate']
```

## Type Definitions

### Stage Types

```typescript
import type { StageId, Stage } from './lever.zod';

// StageId is a union type
type StageId = 'learn' | 'execute' | 'value' | 'enable' | 'replicate';

// Stage interface
interface Stage {
  id: StageId;
  name: string;
  order: number;
  question: string;
  output: string;
  description: string;
  characteristics?: string[];
}
```

### Framework Types

```typescript
import type { FrameworkId, Framework, FrameworkStage } from './lever.zod';

type FrameworkId =
  | 'bloom'
  | 'dreyfus'
  | 'shu-ha-ri'
  | 'apprenticeship'
  | 'conscious-competence';

interface Framework {
  id: FrameworkId;
  name: string;
  description: string;
  origin: string;
  primary_focus: string;
  stages: FrameworkStage[];
}
```

### Profile Types

```typescript
import type { Profile, StageInterpretation, Competency } from './lever.zod';

interface Profile {
  id: string;
  name: string;
  description: string;
  domain: string;
  stage_interpretations: StageInterpretation[];
}

interface StageInterpretation {
  stage_id: StageId;
  title?: string;
  description: string;
  competencies?: Competency[];
  evidence?: string[];
  examples?: string[];
}
```

## Available Schemas

| Schema | Description |
|--------|-------------|
| `stageIdSchema` | LEVER stage identifier enum |
| `stageSchema` | Complete stage definition |
| `frameworkIdSchema` | Classical framework identifier enum |
| `frameworkSchema` | Classical framework definition |
| `stageMappingSchema` | LEVER to framework stage mapping |
| `mappingTableSchema` | Full mapping matrix |
| `profileSchema` | Domain-specific profile |
| `specSchema` | Complete LEVER specification |

## Constants

```typescript
import {
  STAGES,
  FRAMEWORKS,
  STAGE_PROGRESSION,
  OUTPUT_PROGRESSION
} from './lever.zod';

// All stage IDs in order
STAGES // ['learn', 'execute', 'value', 'enable', 'replicate']

// All framework IDs
FRAMEWORKS // ['bloom', 'dreyfus', 'shu-ha-ri', 'apprenticeship', 'conscious-competence']

// Stage categories
STAGE_PROGRESSION.personal  // ['learn', 'execute', 'value']
STAGE_PROGRESSION.leveraged // ['enable', 'replicate']

// Output by stage
OUTPUT_PROGRESSION.learn     // 'Knowledge'
OUTPUT_PROGRESSION.execute   // 'Work'
OUTPUT_PROGRESSION.value     // 'Outcomes'
OUTPUT_PROGRESSION.enable    // 'People'
OUTPUT_PROGRESSION.replicate // 'Systems'
```

## Validation Examples

### Validating User Input

```typescript
import { stageIdSchema } from './lever.zod';

function setUserStage(input: unknown) {
  const result = stageIdSchema.safeParse(input);
  if (!result.success) {
    console.error('Invalid stage:', result.error.issues);
    return null;
  }
  return result.data;
}

setUserStage('value');   // Returns 'value'
setUserStage('invalid'); // Returns null, logs error
```

### Validating API Responses

```typescript
import { specSchema } from './lever.zod';

async function fetchLeverSpec(): Promise<Spec | null> {
  const response = await fetch('/api/lever');
  const data = await response.json();

  const result = specSchema.safeParse(data);
  if (!result.success) {
    console.error('Invalid spec:', result.error.format());
    return null;
  }

  return result.data;
}
```

### Creating Custom Profiles

```typescript
import { profileSchema, type Profile } from './lever.zod';

const myProfile: Profile = {
  id: 'data-scientist',
  name: 'Data Scientist',
  description: 'LEVER for data science roles',
  domain: 'data',
  stage_interpretations: [
    {
      stage_id: 'learn',
      title: 'Analyst',
      description: 'Learning data fundamentals',
      evidence: ['Completes statistics courses', 'Learns Python/R'],
    },
    // ... other stages
  ],
};

// Validate the profile
const validated = profileSchema.parse(myProfile);
```

## Pure TypeScript (No Zod)

If you don't need runtime validation, use the pure TypeScript interfaces:

```typescript
import type {
  StageId,
  Stage,
  Framework,
  Profile,
  Spec
} from './lever.d';

// Type-safe without Zod dependency
function processStage(stage: Stage) {
  console.log(`${stage.order}. ${stage.name}`);
}
```

## Type Synchronization

LEVER types are maintained in sync with Go types:

1. **Go structs** are the source of truth (`pkg/lever/`)
2. **JSON Schema** is generated from Go (`schema/`)
3. **TypeScript/Zod** is manually synchronized (`ts/`)

When Go types change, update TypeScript files to match.
