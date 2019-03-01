package flowascode

import yaml "github.com/go-yaml/yaml"

// NewFromYAML takes a yaml string and creates a Flow object
func NewFromYAML(yaml string) (*Flow, error) {
	var flow Flow
	err := yaml.Unmarshal([]byte(yaml), &flow)
	return &flow, err
}
