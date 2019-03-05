package flowascode

// HasDefault returns true if there is a default step
//
// A default step gets executed when no step is defined
func (f *Flow) HasDefault() bool {
	return f.HasStep("default")
}
