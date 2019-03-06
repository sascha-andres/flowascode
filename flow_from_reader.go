// Copyright © 2019 Sascha Andres <sascha.andres@outlook.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flowascode

import (
	"errors"
	"github.com/go-yaml/yaml"
	"io"
	"io/ioutil"
)

// NewFromReader creates a Flow from a reader
func NewFromReader(reader io.Reader) (*Flow, error) {
	log := logger.WithField("method", "NewFromReader")

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
