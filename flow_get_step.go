package flowascode

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// GetStep returns the named step if found
func (f *Flow) GetStep(name string) (*Step, error) {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "GetStep")

	log.Debugf("called with name := [%s]", name)

	for _, value := range f.Steps {
		if strings.ToLower(value.Name) == strings.ToLower(name) {
			return &value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("found no [%s] step", name))
}
