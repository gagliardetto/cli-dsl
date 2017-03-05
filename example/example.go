package main

import (
	"fmt"
	"regexp"

	. "github.com/gagliardetto/cli-dsl"
)

func main() {
	all := Root("some-cli",
		func() {
			Flags(func() {
				Flag("debug", Int, func() {
					//Short("i")
					//Default(12)
					Required()
					//Usage("some explanation")
					//Nameless() //meaning that it can be specified as: `use some_id` or `use id=some_id`

					MinValue(1)
				})
			})
		}, func() {

			Command(
				"use",
				func() {
					Description("Short description.", "A longer description.")
					Examples("use 1", "use 2")
					Run(func(ctx *Ctx) {
						fmt.Println("this is the 'use' command")
						return
					})

					Flags(func() {
						Flag("name", Int, func() {
							//Short("i")
							//Default(12)
							Required()
							//Usage("some explanation")
							//Nameless() //meaning that it can be specified as: `use some_id` or `use id=some_id`

							MinValue(1)
						})
						Flag("instance-id", String, func() {
							//Short("I")
							MinLength(1)
							MustRegex(regexp.MustCompile("i-([a-z0-9]+)"))
						})
					})

				},
			)

		})

	fmt.Println(all.Execute())
}
