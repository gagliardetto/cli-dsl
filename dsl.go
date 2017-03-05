package cli

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-stack/stack"
	"github.com/spf13/cobra"
)

// currentDefinition is the current working element
var currentDefinition interface{}

// Root defines a new Root
func Root(use string, Is ...I) *CommandStruct {
	newRouter := New()
	newRouter.Use = use
	currentDefinition = newRouter

	for _, i := range Is {
		switch t := i.(type) {
		case func():
			{
				i.(func())()
			}
		default:
			pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
		}
	}

	return newRouter
}

// Command declares a new command in the CommRouter
func Command(use string, Is ...I) {

	previousDefinition := currentDefinition
	defer func() {
		currentDefinition = previousDefinition
	}()

	switch t := currentDefinition.(type) {
	case *CommandStruct:
		{
			rr := currentDefinition.(*CommandStruct)
			subCommand := New()

			subCommand.Use = use

			currentDefinition = subCommand
			for _, i := range Is {
				switch t := i.(type) {
				case func():
					{
						i.(func())()
					}
				default:
					pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
				}
			}

			rr.AddCommand(&subCommand.Command)
		}
	default:
		{
			pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
		}
	}

}

// Description sets the description
func Description(short, long string) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		currentDefinition.(*CommandStruct).Command.Short = short
		currentDefinition.(*CommandStruct).Command.Long = long
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// Examples sets the examples
func Examples(examples ...string) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		currentDefinition.(*CommandStruct).Command.Example = strings.Join(examples, "\n")
		fmt.Println("setting Examples for", currentDefinition.(*CommandStruct).Command.Use)
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// Run sets the examples
func Run(f func(ctx *Ctx)) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		currentDefinition.(*CommandStruct).Command.Run = func() func(cmd *cobra.Command, args []string) {
			fmt.Println("setting Run for", currentDefinition.(*CommandStruct).Command.Use)
			rr := func(cmd *cobra.Command, args []string) {
				// parse and validate stuff
				// create Ctx
				// pass Ctx to f
				ctx := &Ctx{}
				f(ctx)
			}
			return rr
		}()
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// Flags defines a list of flags and their requirements
func Flags(Is ...I) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		//cmd := currentDefinition.(*CommandStruct)
		fmt.Println("setting Flags for", currentDefinition.(*CommandStruct).Command.Use)

		for _, i := range Is {
			switch t := i.(type) {
			case func():
				{
					i.(func())()
				}
			default:
				pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
			}
		}

	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// Flag is a single flag with its requirements
func Flag(key string, flagType Type, Is ...I) {
	previousDefinition := currentDefinition
	defer func() {
		currentDefinition = previousDefinition
	}()

	switch t := currentDefinition.(type) {
	case *CommandStruct:
		cmd := currentDefinition.(*CommandStruct)
		fmt.Println("setting Flag for", currentDefinition.(*CommandStruct).Command.Use, fmt.Sprintf("%q", key))

		//req := &CommandStruct{}
		//req.Type = flagType
		//currentDefinition = req

		// debug
		a := ""
		cmd.Flags().StringVar(&a, key, "flag default value", "flag usage")

		for _, i := range Is {
			switch t := i.(type) {
			case func():
				{
					i.(func())()
				}
			default:
				pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
			}
		}

		//cmd.AddParamRequirement(key, currentDefinition.(*CommandStruct))

	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// Required tells that this field is required
func Required() {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		req := currentDefinition.(*CommandStruct)
		req.Required = true
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// MinValue sets the min value for the field
func MinValue(val interface{}) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		req := currentDefinition.(*CommandStruct)
		req.MinValue = val
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// MinLength sets the min length
func MinLength(length int) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		req := currentDefinition.(*CommandStruct)
		req.MinLength = &length
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

// MustRegex sets the regexp that the value must match
func MustRegex(regex *regexp.Regexp) {
	switch t := currentDefinition.(type) {
	case *CommandStruct:
		req := currentDefinition.(*CommandStruct)
		req.MustRegex = regex
	default:
		pnk(fmt.Sprintf("Unexpected type: \n %#v", t))
	}
}

func pnk(msg string) {
	// s := stack.Trace().TrimBelow(stack.Caller(2)).TrimRuntime()
	s := stack.Trace().TrimRuntime()
	if len(s) > 0 {
		panic(msg + fmt.Sprintf("\n\n%v   %[1]n()", s))
	}
}
