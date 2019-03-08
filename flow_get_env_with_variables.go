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
	"fmt"
	"os"
)

// getEnvWithVariables returns a list of environment variable amended with the
// variables passed to it
func (f *Flow) getEnvWithVariables(step *Step, variables map[string]string) []string {
	log := logger.WithField("method", "*Flow.getEnvWithVariables")

	log.Debugf("called with step := [%s]", step)
	log.Debugf("called with f := [%s]", f)
	log.Debugf("called with variables := [%s]", variables)

	if nil == variables {
		return nil
	}
	environmentVariables := os.Environ()
	for key, value := range variables {
		environmentVariables = append(environmentVariables,
			fmt.Sprintf("%s=%s",
				key,
				step.applyTextTemplate(f, os.ExpandEnv(value))))
	}
	return environmentVariables
}
