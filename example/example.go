package main

import (
	"os"

	"github.com/seeruk/go-console"
	"github.com/seeruk/go-console/parameters"
)

func main() {
	application := console.NewApplication("seeruk/go-console", "0.1.0")
	application.Logo = `
                                             #
                              ###            ##
######## ### ####### #######  ###   #######  ###  ##
         ###       ##      ## ###         ## #### ##
 ####### ###  ###  ## ##   ## ###    ##   ## #######
 ###     ###  ###  ## ##   ## ###    ##   ## ### ###
 ####### ###  ######   #####  ####### #####  ###  ##
                                                   #
`

	var isMarmiteNice bool
	var isVerbose bool
	var name = "World"
	var favNum int

	application.AddGlobalOption(console.OptionDefinition{
		Value:  parameters.NewBoolValue(&isMarmiteNice),
		Spec:   "-m, --marmite",
		Desc:   "Is marmite nice?",
		EnvVar: "MARMITE_NICE",
	})

	application.AddCommand(&console.Command{
		Name:        "greet",
		Alias:       "g",
		Description: "Greet's the given user, or the world.",
		Help:        "You don't have to specify a name.",
		Configure: func(definition *console.Definition) {
			definition.AddOption(console.OptionDefinition{
				Value:  parameters.NewBoolValue(&isVerbose),
				Spec:   "-v, --isVerbose",
				Desc:   "Is isVerbose mode enabled?",
				EnvVar: "EXAMPLE_VERBOSE",
			})

			definition.AddOption(console.OptionDefinition{
				Value:  parameters.NewStringValue(&name),
				Spec:   "-n, --name=NAME",
				Desc:   "Provide a name for the greeting.",
				EnvVar: "EXAMPLE_NAME",
			})

			definition.AddArgument(console.ArgumentDefinition{
				Value: parameters.NewIntValue(&favNum),
				Spec:  "FAVOURITE_NUMBER",
				Desc:  "Provide your favourite number.",
			})
		},
		Execute: func(input *console.Input, output *console.Output) error {
			output.Printf("Hello, %s!\n", name)
			output.Printf("Your favourite number is %d.\n", favNum)
			output.Printf("Is isVerbose mode enabled? %t\n", isVerbose)
			output.Printf("Oh, by the way. Is marmite nice? %t\n", isMarmiteNice)
			return nil
		},
	})

	var test string
	var test2 string

	application.SetRootCommand(&console.Command{
		Configure: func(definition *console.Definition) {
			definition.AddOption(console.OptionDefinition{
				Value: parameters.NewStringValue(&test),
				Spec:  "--test=TEST",
				Desc:  "Test option for root command",
			})

			definition.AddArgument(console.ArgumentDefinition{
				Value: parameters.NewStringValue(&test2),
				Spec:  "TEST2",
				Desc:  "Test argument for root command",
			})
		},
		Execute: func(input *console.Input, output *console.Output) error {
			output.Println("Hello, World!")
			return nil
		},
	})

	code := application.Run(os.Args[1:], os.Environ())

	os.Exit(code)
}
