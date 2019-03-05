package flowascode

// GetDefault returns the default step
func (f *Flow) GetDefault() (*Step, error) {
	return f.GetStep("default")
}
