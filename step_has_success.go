package flowascode

// HasSuccess returns true if there are further steps
// in case of a successful run defined
func (s *Step) HasSuccess() bool {
	if len(s.OnSuccess) > 0 {
		return true
	}
	return false
}
