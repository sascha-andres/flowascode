package flowascode

// HasFailure returns true if there are further steps
// in case of a non successful run defined
func (s *Step) HasFailure() bool {
	if len(s.OnFailure) > 0 {
		return true
	}
	return false
}
