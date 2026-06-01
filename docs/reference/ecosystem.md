# LEVER Ecosystem

LEVER is designed to integrate with other capability and maturity frameworks.

## Related Frameworks

### PBMM (Product Builder Maturity Model)

**Product Builder Maturity Model** provides detailed assessment criteria for Product Builders at each LEVER stage.

| LEVER Stage | PBMM Level | Focus |
|-------------|------------|-------|
| Learn | Associate | Learning fundamentals |
| Execute | PB I | Shipping with guidance |
| Value | PB II | Independent delivery |
| Enable | Senior PB | Team development |
| Replicate | Principal PB | System creation |

PBMM extends the Product Builder profile with:

- Detailed competency matrices
- Assessment rubrics
- Career progression guides
- Development recommendations

**Learn more:** [Product Builders HQ](https://productbuildershq.com)

---

### ASDM (Autonomous Software Delivery Model)

**Autonomous Software Delivery Model** applies LEVER concepts to team capability rather than individual progression.

| LEVER Stage | ASDM Level | Team Capability |
|-------------|------------|-----------------|
| Learn | Level 1 | Learning processes |
| Execute | Level 2 | Following playbooks |
| Value | Level 3 | Autonomous delivery |
| Enable | Level 4 | Cross-team enablement |
| Replicate | Level 5 | Platform/system creation |

ASDM provides:

- Team maturity assessment
- Capability benchmarks
- Transformation roadmaps
- Metrics and measurement

**Learn more:** [Product Builders HQ](https://productbuildershq.com)

---

### PRISM (AI Integration)

**PRISM** integrates AI capabilities with human capability progression.

LEVER + AI acceleration:

| LEVER Stage | AI Amplification |
|-------------|------------------|
| Learn | AI tutoring, personalized learning |
| Execute | AI coding assistants, automation |
| Value | AI-augmented decision making |
| Enable | AI-powered coaching tools |
| Replicate | AI agents as multipliers |

---

## Integration Patterns

### Using LEVER as Foundation

```
┌─────────────────────────────────────────────────┐
│                    LEVER                         │
│  (Universal Capability Framework)               │
├─────────────────┬─────────────────┬─────────────┤
│      PBMM       │      ASDM       │   Custom    │
│  (Individuals)  │    (Teams)      │  Profiles   │
└─────────────────┴─────────────────┴─────────────┘
```

### Profile Extension

Organizations can create domain-specific profiles that:

1. **Map to LEVER stages** — Use the 5-stage progression
2. **Add domain competencies** — Define specific capabilities
3. **Integrate with tools** — Connect to existing HR/development systems
4. **Measure progress** — Use LEVER questions as assessment criteria

### Maturity Model Pattern

LEVER stages map naturally to maturity levels:

```go
// Map LEVER to your maturity model
var maturityMapping = map[lever.StageID]int{
    lever.StageLearn:     1,
    lever.StageExecute:   2,
    lever.StageValue:     3,
    lever.StageEnable:    4,
    lever.StageReplicate: 5,
}
```

## Adoption Guide

### For Organizations

1. **Start with reference profiles** — Use Product Builder or Engineer profiles
2. **Customize for your domain** — Create profiles that match your roles
3. **Integrate with career ladders** — Map LEVER stages to your levels
4. **Build assessment tools** — Use the stage questions for evaluation

### For Tool Builders

1. **Consume the schema** — Use JSON Schema for validation
2. **Import the types** — Use Go or TypeScript types
3. **Build on mappings** — Leverage classical framework mappings
4. **Extend with profiles** — Create domain-specific profiles

### For Researchers

1. **Reference the mappings** — Cite LEVER's synthesis of classical frameworks
2. **Study the progression** — Enable/Replicate as unique contributions
3. **Apply to new domains** — Create profiles for unexplored areas

## Resources

- **GitHub:** [github.com/ProductBuildersHQ/lever](https://github.com/ProductBuildersHQ/lever)
- **Product Builders HQ:** [productbuildershq.com](https://productbuildershq.com)
- **JSON Schemas:** [Schema Reference](schema.md)

## Contributing

LEVER welcomes:

- **New profiles** — Domain-specific interpretations
- **Improved mappings** — Refinements to framework mappings
- **Tool integrations** — Libraries and utilities
- **Documentation** — Examples and guides

See the GitHub repository for contribution guidelines.
