package flowascode

const ErrNoScript = Error("no script in step")

// Execute is going to execute the script
// calls os specific implementations
func (s *Step) Execute(pathToShell string, variables []string) error {
	log := logger.WithField("method", "*Step.Execute")

	log.Debugf("called with pathToShell := [%s]", pathToShell)
	log.Debugf("called with variables := %v", variables)

	if len(s.Script) == 0 {
		return ErrNoScript
	}

	return s.executeOS(pathToShell, variables)
}
