package flowascode

import "github.com/sirupsen/logrus"

// HasDefault returns true if there is a default step
//
// A default step gets executed when no step is defined
func (f *Flow) HasDefault() bool {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "HasDefault")

	log.Debug("called")

	return f.HasStep("default")
}
