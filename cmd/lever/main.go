// Command lever provides CLI tools for the LEVER Capability Framework.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ProductBuildersHQ/lever/pkg/lever"
	"github.com/invopop/jsonschema"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "schema":
		if err := generateSchema(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "spec":
		if err := printSpec(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "profile-schema":
		if err := generateProfileSchema(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "stages":
		printStages()
	case "mappings":
		if err := printMappings(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "version":
		fmt.Println(lever.Version)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`lever - LEVER Capability Framework CLI

Usage:
  lever <command>

Commands:
  schema          Generate JSON Schema for lever.Spec (stdout)
  profile-schema  Generate JSON Schema for lever.Profile (stdout)
  spec            Output the complete LEVER spec as JSON (stdout)
  stages          List LEVER stages
  mappings        Output the mapping table as JSON
  version         Print the spec version`)
}

func generateSchema() error {
	r := new(jsonschema.Reflector)
	r.DoNotReference = false

	schema := r.Reflect(&lever.Spec{})
	schema.ID = "https://productbuildershq.com/schemas/lever.schema.json"
	schema.Title = "LEVER Capability Framework"
	schema.Description = lever.NewSpec().Description

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(schema)
}

func generateProfileSchema() error {
	r := new(jsonschema.Reflector)
	r.DoNotReference = false

	schema := r.Reflect(&lever.Profile{})
	schema.ID = "https://productbuildershq.com/schemas/profile.schema.json"
	schema.Title = "LEVER Profile"
	schema.Description = "A domain-specific interpretation of LEVER stages"

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(schema)
}

func printSpec() error {
	spec := lever.NewSpec()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(spec)
}

func printStages() {
	for _, s := range lever.Stages() {
		fmt.Printf("%d. %s (%s)\n", s.Order, s.Name, s.ID)
		fmt.Printf("   Question: %s\n", s.Question)
		fmt.Printf("   Output:   %s\n", s.Output)
		fmt.Println()
	}
}

func printMappings() error {
	table := lever.BuildMappingTable()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(table)
}
