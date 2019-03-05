package flowascode

// GetDefault returns the default step
func (f *Flow) GetDefault() (*Step, error) {
	log := logger.WithField("method", "GetDefault")

	log.Debug("called")

	return f.GetStep("default")
}
