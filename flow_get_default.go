package flowascode

import "github.com/sirupsen/logrus"

// GetDefault returns the default step
func (f *Flow) GetDefault() (*Step, error) {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "GetDefault")

	log.Debug("called")

	return f.GetStep("default")
}
