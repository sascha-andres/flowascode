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

// Execute takes a step name and starts to execute from there
// if no ma,e is given it tries to get the default namespace
func (f *Flow) Execute(name string, variables map[string]string) error {
	log := logger.WithField("method", "*Flow.Execute")

	log.Debugf("called with name := [%s]", name)
	log.Debugf("called with variables := %v", variables)

	var (
		step *Step
		err  error
	)
	if name == "" && f.HasDefault() {
		step, err = f.GetDefault()
	} else {
		step, err = f.GetStep(name)
	}
	if err != nil {
		log.Errorf("error finding step: %s", err)
		return err
	}
	err = f.ValidateStep(name)
	if err != nil {
		log.Errorf("step not valid: %s", err)
		return err
	}
	err = step.Execute(f.Shell, getEnvWithVariables(variables))
	if err != nil {
		log.Errorf("step execution failed: %s", err)
		for _, value := range step.OnFailure {
			_ = f.Execute(value.Name, value.Variables)
		}
	} else {
		log.Infof("step execution succeeded: [%s]", step)
		for _, value := range step.OnSuccess {
			_ = f.Execute(value.Name, value.Variables)
		}
	}
	return err
}

// getEnvWithVariables returns a list of environment variable amended with the
// variables passed to it
func getEnvWithVariables(variables map[string]string) []string {
	if nil == variables {
		return nil
	}
	environmentVariables := os.Environ()
	for key, value := range variables {
		environmentVariables = append(environmentVariables, fmt.Sprintf("%s=%s", key, value))
	}
	return environmentVariables
}
