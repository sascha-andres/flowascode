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

import "github.com/sirupsen/logrus"

type (

	// Error type for constant error definitions
	Error string

	// Flow represents a complete flow
	Flow struct {
		// BreakOnError bool   `yaml:"break_on_error"` // BreakOnError breaks if a script returns a non zero result
		Name        string `yaml:"name"`        // Name is a name for the flow
		Description string `yaml:"description"` // Description is a more descriptive text what the flow does
		Steps       []Step `yaml:"steps"`       // Steps is the collection of steps in the flow
		Shell       string `yaml:"shell"`       // Shell is the path or binary name for the shell to use to execute the script
	}

	// Step is a single execution point in a flow
	Step struct {
		Name      string       `yaml:"name"`       // Name of a step, also used to reference from C(Descendant)
		Script    []string     `yaml:"script"`     // Script to execute
		OnSuccess []Descendant `yaml:"on_success"` // OnSuccess is a list of steps to execute if the script returns with 0
		OnFailure []Descendant `yaml:"on_failure"` // OnFailure is a list of steps to execute if the scripts returns not a 0
	}

	// Descendant represents on step after the current step
	// that is it is a reference to a number of other steps
	Descendant struct {
		Name      string            `yaml:"name"`      // Name is the referenced step
		Variables map[string]string `yaml:"variables"` // Variables is a list of variables that are set as environment variables
	}
)

var (
	logger *logrus.Entry
)

func (e Error) Error() string { return string(e) }

func init() {
	logrus.SetLevel(logrus.WarnLevel)
	logger = logrus.WithField("package", "flowascode")
}

// SetLogLevel can be used to adjust the log level
func SetLogLevel(level logrus.Level) {
	logrus.SetLevel(level)
	logger = logrus.WithField("package", "flowascode")
}
