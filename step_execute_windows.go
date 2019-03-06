package flowascode

// executeOS implementation for Windows
func (s *Step) executeOS(pathToShell string, variables []string) error {
	log := logger.WithField("method", "*Step.executeOS[windows]")

	log.Debugf("called with pathToShell := [%s]", pathToShell)

	return errors.New("not yet implemented")
}
