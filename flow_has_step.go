package flowascode

import "strings"

// HasStep returns true in case the step exists
func (f *Flow) HasStep(name string) bool {
	for _, value := range f.Steps {
		if strings.ToLower(value.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}
