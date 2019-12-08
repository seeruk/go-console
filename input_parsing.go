package console

import (
	"fmt"
	"strings"

	"github.com/seeruk/go-console/parameters"
)

// ParseInput takes the raw input, and regardless of what is actually defined in the definition,
// categorises the input as either arguments or options. In other words, the raw input is iterated
// over, not the definition's parameters. The definition is used so that we can identify options
// that should have values and consume the next argument as it's value.
func ParseInput(definition *Definition, args []string) *Input {
	var input Input
	var optsEnded bool

	// We don't range, because we can modify `i` in the middle of the loop this way. This allows us
	// to consume the next argument if we want (and if it's available).
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--" {
			optsEnded = true
			continue
		}

		argLen := len(arg)
		isLongOpt := argLen > 2 && strings.HasPrefix(arg, "--")
		isShortOpt := argLen > 1 && strings.HasPrefix(arg, "-")

		if !optsEnded && (isLongOpt || isShortOpt) {
			var options []InputOption

			if isLongOpt {
				options = parseOption(arg, "--")
			} else {
				options = parseOption(arg, "-")
			}

			// mappedOptions is a temporary place for the options parsed from this one argument to
			// be stored. We need this so we can always identify if we actually parsed any options
			// at all, and so we can get the last option for this argument all of the time.
			mappedOptions := []InputOption{}

			for _, option := range options {
				// All options will be mapped, regardless. Only the last option will have any value.
				mappedOptions = append(mappedOptions, option)
			}

			mappedOptionsLen := len(mappedOptions)

			if mappedOptionsLen > 0 {
				lastOption := mappedOptions[mappedOptionsLen-1]

				defOpt, exists := definition.options[lastOption.Name]
				if !exists {
					// We don't care about options that don't exist in the definition. We shouldn't
					// be consuming arguments for them, because they won't require a value.
					break
				}

				isRequired := defOpt.ValueMode == parameters.OptionValueRequired
				hasArgsLeft := len(args) > (i + 1) // Length required for next is +2, not +1.
				hasNoValYet := lastOption.Value == ""

				// If the value is required, but we don't yet have a value on the option, this means
				// we'll consume the next argument following the option and treat it as the value.
				if isRequired && hasNoValYet && hasArgsLeft {
					lastOption.Value = args[i+1]
					i++
				}

				mappedOptions[mappedOptionsLen-1] = lastOption
			}

			input.Options = append(input.Options, mappedOptions...)
		} else {
			input.Arguments = append(input.Arguments, InputArgument{Value: arg})
		}
	}

	return &input
}

// parseOption parses an input option with the given prefix (e.g. '-', or '--'). It returns an array
// because short options can contain multiple options without values.
func parseOption(option string, prefix string) []InputOption {
	var results []InputOption

	trimmed := strings.TrimPrefix(option, prefix)
	split := strings.SplitN(trimmed, "=", 2)

	var key string
	var val string

	if len(split) >= 1 {
		key = split[0]
	}

	if len(split) == 2 {
		val = split[1]
	}

	if prefix == "-" {
		results = append(results, parseShortOption(key, val)...)
	}

	if prefix == "--" {
		results = append(results, parseLongOption(key, val))
	}

	return results
}

func parseLongOption(key, value string) InputOption {
	return InputOption{Name: key, Value: value}
}

func parseShortOption(key, value string) []InputOption {
	var results []InputOption

	// Convert key into rune slice, so we can iterate over each rune properly.
	runes := []rune(key)

	// Folded options
	if len(key) > 1 {
		// We want to handle the last run differently, so that the value following all of the
		// options can be given to the last option.
		for i := 0; i < len(runes)-1; i++ {
			results = append(results, InputOption{Name: fmt.Sprintf("%c", runes[i])})
		}

		results = append(results, InputOption{Name: fmt.Sprintf("%c", runes[len(runes)-1]), Value: value})
	} else {
		results = append(results, InputOption{Name: key, Value: value})
	}

	return results
}
