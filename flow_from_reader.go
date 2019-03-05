package flowascode

import (
	"errors"
	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

// NewFromReader creates a Flow from a reader
func NewFromReader(reader io.Reader) (*Flow, error) {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "NewFromReader")

	log.Debugf("called")

	if nil == reader {
		return nil, errors.New("nil reader not allowed")
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var flow Flow
	err = yaml.Unmarshal(data, &flow)
	return &flow, err
}
