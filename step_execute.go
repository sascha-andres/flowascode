package flowascode

import "github.com/sirupsen/logrus"

const ErrNoScript = Error("no script in step")

// Execute is going to execute the script
// calls os specific implementations
func (s *Step) Execute(pathToShell string) error {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "Execute")

	log.Debugf("called with pathToShell := [%s]", pathToShell)

	if len(s.Script) == 0 {
		return ErrNoScript
	}

	return s.executeOS(pathToShell)
}
