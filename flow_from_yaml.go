package flowascode

import yaml "gopkg.in/yaml.v2"

func NewFromYAML(yaml string) (*Flow, error) {
	var flow Flow
	err := yaml.Unmarshal([]byte(yaml), &flow)
	return &flow, err
}
