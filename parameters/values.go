package parameters

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"
)

// Value contains the values stored in a Parameter. This interface has an implementation for each
// type of value that may be accepted.
//
// Set is used to assign the values upon parsing and validation. All input comes in as a string.
type Value interface {
	Set(string) error
	String() string
}

// FlagValue represents an option that does not have a value. Allows toggling behaviour.
type FlagValue interface {
	Value
	FlagValue() string
}

// BoolValue abstracts functionality for parsing input that should be represented as a boolean. The
// BoolValue type also implements the FlagValue interface so that an alternative to the default
// value can be used if no value is present.
type BoolValue bool

// NewBoolValue creates a new BoolValue.
func NewBoolValue(ref *bool) *BoolValue {
	return (*BoolValue)(ref)
}

// Set assigns a value to the value that this BoolValue references.
func (v *BoolValue) Set(s string) error {
	b, err := strconv.ParseBool(s)
	*v = BoolValue(b)
	return err
}

// String converts this BoolValue to a string.
func (v *BoolValue) String() string {
	return fmt.Sprintf("%v", *v)
}

// FlagValue returns the default value boolValue when no value is present (i.e. when used as a flag)
func (v *BoolValue) FlagValue() string {
	return "true"
}

// DateValue abstracts functionality for parsing input that should be represented as a time.Time.
type DateValue time.Time

// NewDateValue creates a new DateValue.
func NewDateValue(ref *time.Time) *DateValue {
	return (*DateValue)(ref)
}

// Set assigns a value to the value that this DateValue references.
func (v *DateValue) Set(s string) error {
	d, err := time.Parse("2006-01-02", s)
	*v = DateValue(d)
	return err
}

// String converts this DateValue to a string.
func (v *DateValue) String() string {
	return (*time.Time)(v).Format("2006-01-02")
}

// DurationValue abstracts functionality for parsing input that should be represented as a
// time.Duration.
type DurationValue time.Duration

// NewDurationValue creates a new DurationValue.
func NewDurationValue(ref *time.Duration) *DurationValue {
	return (*DurationValue)(ref)
}

// Set assigns a value to the value that this DurationValue references.
func (v *DurationValue) Set(s string) error {
	d, err := time.ParseDuration(s)
	*v = DurationValue(d)
	return err
}

// String converts this DurationValue to a string.
func (v *DurationValue) String() string {
	return (*time.Duration)(v).String()
}

// Float32Value abstracts functionality for parsing input that should be represented as a float32.
type Float32Value float32

// NewFloat32Value creates a new Float32Value.
func NewFloat32Value(ref *float32) *Float32Value {
	return (*Float32Value)(ref)
}

// Set assigns a value to the value that this Float32Value references.
func (v *Float32Value) Set(s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}

	*v = Float32Value(float32(f))
	return err
}

// String converts this float32Value to a string.
func (v *Float32Value) String() string {
	return fmt.Sprintf("%v", *v)
}

// Float64Value abstracts functionality for parsing input that should be represented as a float64.
type Float64Value float64

// NewFloat64Value creates a new float64Value.
func NewFloat64Value(ref *float64) *Float64Value {
	return (*Float64Value)(ref)
}

// Set assigns a value to the value that this Float64Value references.
func (v *Float64Value) Set(s string) error {
	f, err := strconv.ParseFloat(s, 64)
	*v = Float64Value(f)
	return err
}

// String converts this Float64Value to a string.
func (v *Float64Value) String() string {
	return fmt.Sprintf("%v", *v)
}

// IntValue abstracts functionality for parsing input that should be represented as an int.
type IntValue int

// NewIntValue creates a new intValue.
func NewIntValue(ref *int) *IntValue {
	return (*IntValue)(ref)
}

// Set assigns a value to the value that this IntValue references.
func (v *IntValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 64)
	*v = IntValue(i)
	return err
}

// String converts this IntValue to a string.
func (v *IntValue) String() string {
	return fmt.Sprintf("%v", *v)
}

// IPValue abstracts functionality for parsing input that should be represented as an IP address.
type IPValue net.IP

// NewIPValue creates a new IPValue.
func NewIPValue(ref *net.IP) *IPValue {
	return (*IPValue)(ref)
}

// Set assigns a value to the value that this IPValue references.
func (v *IPValue) Set(val string) error {
	ip := net.ParseIP(val)
	if ip == nil {
		return fmt.Errorf("invalid IP address format '%v'", val)
	}

	*v = IPValue(ip)

	return nil
}

// String converts this IPValue to a string.
func (v *IPValue) String() string {
	ip := net.IP(*v)

	return ip.String()
}

// StringValue accepts string input, and transparently assigns it to a pointer.
type StringValue string

// NewStringValue creates a new StringValue.
func NewStringValue(ref *string) *StringValue {
	return (*StringValue)(ref)
}

// Set assigns a value to the value that this StringValue references.
func (v *StringValue) Set(val string) error {
	*v = StringValue(val)
	return nil
}

// String converts this StringValue to a string.
func (v *StringValue) String() string {
	return fmt.Sprintf("%v", *v)
}

// URLValue abstracts functionality for parsing input that should be represented as a URL.
type URLValue url.URL

// NewURLValue creates a new URLValue.
func NewURLValue(ref *url.URL) *URLValue {
	return (*URLValue)(ref)
}

// Set assigns a value to the value that this URLValue references.
func (v *URLValue) Set(val string) error {
	res, err := url.Parse(val)
	if err != nil {
		return err
	}
	*v = URLValue(*res)
	return nil
}

// String converts this URLValue to a string.
func (v *URLValue) String() string {
	u := url.URL(*v)

	return fmt.Sprintf("%v", u.String())
}
