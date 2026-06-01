// Generated from Go types in pkg/lever - DO NOT EDIT MANUALLY
// These types match the Zod schemas in lever.zod.ts

// =============================================================================
// Stage Types
// =============================================================================

/** LEVER stage identifiers */
export type StageId = 'learn' | 'execute' | 'value' | 'enable' | 'replicate';

/** A LEVER capability stage */
export interface Stage {
  id: StageId;
  name: string;
  order: number;
  question: string;
  output: string;
  description: string;
  characteristics?: string[];
}

// =============================================================================
// Classical Framework Types
// =============================================================================

/** Classical framework identifiers */
export type FrameworkId =
  | 'bloom'
  | 'dreyfus'
  | 'shu-ha-ri'
  | 'apprenticeship'
  | 'conscious-competence';

/** A stage within a classical framework */
export interface FrameworkStage {
  id: string;
  name: string;
  order: number;
  description: string;
}

/** A classical learning/mastery framework */
export interface Framework {
  id: FrameworkId;
  name: string;
  description: string;
  origin: string;
  primary_focus: string;
  stages: FrameworkStage[];
}

// =============================================================================
// Mapping Types
// =============================================================================

/** Mapping from a LEVER stage to classical framework stages */
export interface StageMapping {
  lever_stage: StageId;
  framework_id: FrameworkId;
  framework_stages: string[];
  notes?: string;
}

/** All mappings for a single classical framework */
export interface FrameworkMappings {
  framework_id: FrameworkId;
  mappings: StageMapping[];
}

/** A row in the mapping table */
export interface MappingRow {
  lever_stage: StageId;
  cells: Record<FrameworkId, string[]>;
}

/** The full mapping table in matrix form */
export interface MappingTable {
  frameworks: FrameworkId[];
  rows: MappingRow[];
}

// =============================================================================
// Profile Types
// =============================================================================

/** A specific capability within a stage */
export interface Competency {
  id: string;
  name: string;
  description: string;
  indicators?: string[];
}

/** Domain-specific meaning for a LEVER stage */
export interface StageInterpretation {
  stage_id: StageId;
  title?: string;
  description: string;
  competencies?: Competency[];
  evidence?: string[];
  examples?: string[];
}

/** A domain-specific interpretation of LEVER stages */
export interface Profile {
  id: string;
  name: string;
  description: string;
  domain: string;
  stage_interpretations: StageInterpretation[];
}

// =============================================================================
// Complete Specification
// =============================================================================

/** The complete LEVER Capability Framework specification */
export interface Spec {
  version: string;
  name: string;
  description: string;
  stages: Stage[];
  frameworks: Framework[];
  mappings: FrameworkMappings[];
  mapping_table: MappingTable;
  reference_profiles?: Profile[];
}

// =============================================================================
// Stage Progression
// =============================================================================

/** Stage progression categories */
export interface StageProgression {
  personal: StageId[];
  leveraged: StageId[];
}

/** Output produced by each stage */
export type OutputProgression = Record<StageId, string>;
