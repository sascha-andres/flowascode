package flowascode

import "github.com/sirupsen/logrus"

// HasFailure returns true if there are further steps
// in case of a non successful run defined
func (s *Step) HasFailure() bool {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "HasFailure")

	log.Debug("called")

	if len(s.OnFailure) > 0 {
		return true
	}
	return false
}
