package parameters

// Option value modes.
const (
	OptionValueNone OptionValueMode = iota
	// @TODO: Remove this. It makes no sense.
	OptionValueOptional
	OptionValueRequired
)

// OptionValueMode represents the different potential requirements of an option's value.
type OptionValueMode int

// Option provides the internal representation of an input option parameter.
type Option struct {
	// The names of this option.
	Names []string
	// The description of this option.
	Description string
	// The name of an environment variable to read an option value from.
	EnvVar string
	// The value that this option references.
	Value Value
	// Does this option take a value? Is it optional, or required?
	ValueMode OptionValueMode
	// The name of the value (shown in contextual help).
	ValueName string
}
