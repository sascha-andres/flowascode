package flowascode

import (
	"github.com/go-yaml/yaml"
)

// NewFromYAML takes a yaml string and creates a Flow object
func NewFromYAML(data string) (*Flow, error) {
	log := logger.WithField("method", "NewFromYAML")

	log.Debug("called")
	var flow Flow
	err := yaml.Unmarshal([]byte(data), &flow)
	return &flow, err
}
