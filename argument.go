package console

// Argument provides the internal representation of an input argument paremeter.
type Argument struct {
	// The name of this argument.
	Name string
	// The description of this argument.
	Description string
	// The value that this argument references.
	Value ParameterValue
	// Is this argument required?
	Required bool
}
