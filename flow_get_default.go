package flowascode

import (
	"errors"
	"strings"
)

// GetDefault returns the default step
func (f *Flow) GetDefault() (*Step, error) {
	for _, value := range f.Steps {
		if strings.ToLower(value.Name) == "default" {
			return &value, nil
		}
	}
	return nil, errors.New("found no default step")
}
