package flowascode

// executeOS implementation for Windows
func (s *Step) executeOS(pathToShell string) error {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "executeOS[windows]")

	log.Debugf("called with pathToShell := [%s]", pathToShell)

	return errors.New("not yet implemented")
}
