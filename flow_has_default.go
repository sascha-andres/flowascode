package flowascode

import (
	"strings"
)

// HasDefault returns true if there is a default step
//
// A default step gets executed when no step is defined
func (f *Flow) HasDefault() bool {
	for _, value := range f.Steps {
		if strings.ToLower(value.Name) == "default" {
			return true
		}
	}
	return false
}
