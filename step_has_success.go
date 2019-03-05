package flowascode

import "github.com/sirupsen/logrus"

// HasSuccess returns true if there are further steps
// in case of a successful run defined
func (s *Step) HasSuccess() bool {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "HasSuccess")

	log.Debug("called")

	if len(s.OnSuccess) > 0 {
		return true
	}
	return false
}
