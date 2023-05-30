package acal

// sourcer represents any struct that can have a source.
//
//go:generate mockery --name=sourcer --case underscore --inpackage
type sourcer interface {
	// From updates the source of this sourcer to the provided Source.
	From(Source)
	// GetSource returns the source of this sourcer.
	GetSource() Source
}

// Source represents where a Value comes from.
type Source string

// Apply applies this Source on all given sourcer.
func (s Source) Apply(targets ...sourcer) {
	for _, target := range targets {
		target.From(s)
	}
}

// String returns the string representation of this Source.
func (s Source) String() string {
	return string(s)
}

const (
	// These are the sources reserved for internal implementation.
	// Clients MUST not declare custom sources with the same names.
	sourceStaticCalculation      Source = "static_calculation"
	sourceProgressiveCalculation Source = "progressive_calculation"
	sourceDependentCalculation   Source = "dependent_calculation"
	sourceRemoteCalculation      Source = "remote_calculation"

	// These are the common sources that the framework provides for
	// all clients to share.
	SourceRequest  Source = "request"
	SourceHardcode Source = "hardcode"
	SourceUnknown  Source = "unknown"
	SourceLogic    Source = "logic"
)
