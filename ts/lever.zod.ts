// Generated from Go types in pkg/lever - DO NOT EDIT MANUALLY
// Regenerate with: npm run generate
import { z } from 'zod';

// =============================================================================
// Stage Types
// =============================================================================

/** LEVER stage identifiers */
export const stageIdSchema = z.enum([
  'learn',
  'execute',
  'value',
  'enable',
  'replicate',
]);
export type StageId = z.infer<typeof stageIdSchema>;

/** A LEVER capability stage */
export const stageSchema = z.object({
  id: stageIdSchema,
  name: z.string(),
  order: z.number().int().min(1).max(5),
  question: z.string(),
  output: z.string(),
  description: z.string(),
  characteristics: z.array(z.string()).optional(),
});
export type Stage = z.infer<typeof stageSchema>;

// =============================================================================
// Classical Framework Types
// =============================================================================

/** Classical framework identifiers */
export const frameworkIdSchema = z.enum([
  'bloom',
  'dreyfus',
  'shu-ha-ri',
  'apprenticeship',
  'conscious-competence',
]);
export type FrameworkId = z.infer<typeof frameworkIdSchema>;

/** A stage within a classical framework */
export const frameworkStageSchema = z.object({
  id: z.string(),
  name: z.string(),
  order: z.number().int(),
  description: z.string(),
});
export type FrameworkStage = z.infer<typeof frameworkStageSchema>;

/** A classical learning/mastery framework */
export const frameworkSchema = z.object({
  id: frameworkIdSchema,
  name: z.string(),
  description: z.string(),
  origin: z.string(),
  primary_focus: z.string(),
  stages: z.array(frameworkStageSchema),
});
export type Framework = z.infer<typeof frameworkSchema>;

// =============================================================================
// Mapping Types
// =============================================================================

/** Mapping from a LEVER stage to classical framework stages */
export const stageMappingSchema = z.object({
  lever_stage: stageIdSchema,
  framework_id: frameworkIdSchema,
  framework_stages: z.array(z.string()),
  notes: z.string().optional(),
});
export type StageMapping = z.infer<typeof stageMappingSchema>;

/** All mappings for a single classical framework */
export const frameworkMappingsSchema = z.object({
  framework_id: frameworkIdSchema,
  mappings: z.array(stageMappingSchema),
});
export type FrameworkMappings = z.infer<typeof frameworkMappingsSchema>;

/** A row in the mapping table */
export const mappingRowSchema = z.object({
  lever_stage: stageIdSchema,
  cells: z.record(frameworkIdSchema, z.array(z.string())),
});
export type MappingRow = z.infer<typeof mappingRowSchema>;

/** The full mapping table in matrix form */
export const mappingTableSchema = z.object({
  frameworks: z.array(frameworkIdSchema),
  rows: z.array(mappingRowSchema),
});
export type MappingTable = z.infer<typeof mappingTableSchema>;

// =============================================================================
// Profile Types
// =============================================================================

/** A specific capability within a stage */
export const competencySchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string(),
  indicators: z.array(z.string()).optional(),
});
export type Competency = z.infer<typeof competencySchema>;

/** Domain-specific meaning for a LEVER stage */
export const stageInterpretationSchema = z.object({
  stage_id: stageIdSchema,
  title: z.string().optional(),
  description: z.string(),
  competencies: z.array(competencySchema).optional(),
  evidence: z.array(z.string()).optional(),
  examples: z.array(z.string()).optional(),
});
export type StageInterpretation = z.infer<typeof stageInterpretationSchema>;

/** A domain-specific interpretation of LEVER stages */
export const profileSchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string(),
  domain: z.string(),
  stage_interpretations: z.array(stageInterpretationSchema),
});
export type Profile = z.infer<typeof profileSchema>;

// =============================================================================
// Complete Specification
// =============================================================================

/** The complete LEVER Capability Framework specification */
export const specSchema = z.object({
  version: z.string(),
  name: z.string(),
  description: z.string(),
  stages: z.array(stageSchema),
  frameworks: z.array(frameworkSchema),
  mappings: z.array(frameworkMappingsSchema),
  mapping_table: mappingTableSchema,
  reference_profiles: z.array(profileSchema).optional(),
});
export type Spec = z.infer<typeof specSchema>;

// =============================================================================
// Stage Progression
// =============================================================================

/** Stage progression categories */
export const stageProgressionSchema = z.object({
  personal: z.array(stageIdSchema),
  leveraged: z.array(stageIdSchema),
});
export type StageProgression = z.infer<typeof stageProgressionSchema>;

/** Output produced by each stage */
export const outputProgressionSchema = z.record(stageIdSchema, z.string());
export type OutputProgression = z.infer<typeof outputProgressionSchema>;

// =============================================================================
// Constants
// =============================================================================

/** All LEVER stages in order */
export const STAGES: readonly StageId[] = [
  'learn',
  'execute',
  'value',
  'enable',
  'replicate',
] as const;

/** Stage progression split */
export const STAGE_PROGRESSION: StageProgression = {
  personal: ['learn', 'execute', 'value'],
  leveraged: ['enable', 'replicate'],
};

/** Output produced by each stage */
export const OUTPUT_PROGRESSION: OutputProgression = {
  learn: 'Knowledge',
  execute: 'Work',
  value: 'Outcomes',
  enable: 'People',
  replicate: 'Systems',
};

/** All framework IDs */
export const FRAMEWORKS: readonly FrameworkId[] = [
  'bloom',
  'dreyfus',
  'shu-ha-ri',
  'apprenticeship',
  'conscious-competence',
] as const;
