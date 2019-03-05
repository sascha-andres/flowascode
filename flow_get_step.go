package flowascode

import (
	"errors"
	"fmt"
	"strings"
)

// GetStep returns the named step if found
func (f *Flow) GetStep(name string) (*Step, error) {
	for _, value := range f.Steps {
		if strings.ToLower(value.Name) == strings.ToLower(name) {
			return &value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("found no [%s] step", name))
}
