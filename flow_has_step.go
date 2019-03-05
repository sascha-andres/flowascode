package flowascode

import (
	"github.com/sirupsen/logrus"
	"strings"
)

// HasStep returns true in case the step exists
func (f *Flow) HasStep(name string) bool {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "HasStep")

	log.Debugf("called with name := [%s]", name)

	for _, value := range f.Steps {
		if strings.ToLower(value.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}
