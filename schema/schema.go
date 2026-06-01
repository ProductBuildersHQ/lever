// Package schema provides embedded JSON Schema for the LEVER specification.
package schema

import (
	_ "embed"
)

// LeverSchema is the JSON Schema for the LEVER specification.
// Generated from Go types using invopop/jsonschema.
//
//go:embed lever.schema.json
var LeverSchema []byte

// ProfileSchema is the JSON Schema for LEVER profiles.
//
//go:embed profile.schema.json
var ProfileSchema []byte
