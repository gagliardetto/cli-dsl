package cli

import "github.com/spf13/cobra"

func New() *CommandStruct {
	return &CommandStruct{}
}

type CommandStruct struct {
	controller         func(interface{}) error
	selectorValidators requirementsMap
	paramValidators    requirementsMap
	Common

	Requirements

	cobra.Command
}

// Common defines common descriptive elements like description, example
type Common struct {
	names    []string
	Examples []string
}

// I is an interface
type I interface{}

// H is a map with string as key and string as value
type H map[string]string

type Ctx struct {
}
