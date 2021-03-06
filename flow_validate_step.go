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
	"github.com/hashicorp/go-multierror"
)

const ErrMissingSuccessStep = Error("success step not defined")
const ErrMissingFailureStep = Error("failure step not defined")
const ErrMissingStep = Error("step not defined")

// ValidateStep looks if a step could run successful
func (f *Flow) ValidateStep(name string) error {
	log := logger.WithField("method", "*Flow.ValidateStep")

	log.Debugf("called with name := [%s]", name)

	var result *multierror.Error

	step, err := f.validateStepExist(name)
	if err != nil {
		return err
	}
	if err := f.validateSuccess(step); err != nil {
		result = multierror.Append(result, err)
	}
	if err := f.validateFailure(step); err != nil {
		result = multierror.Append(result, err)
	}
	if 0 == len(step.Script) {
		result = multierror.Append(result, ErrNoScript)
	}

	return result.ErrorOrNil()
}

// validateStepExist checks if a step exists
func (f *Flow) validateStepExist(name string) (*Step, error) {
	log := logger.WithField("method", "*Flow.validateStepExist")

	log.Debugf("called with name := [%s]", name)

	var step *Step
	for _, value := range f.Steps {
		if value.Name == name {
			step = &value
			break
		}
	}
	if nil == step {
		return nil, ErrMissingStep
	}
	return step, nil
}

// validateSuccess checks if all required steps on success
// exist
func (f *Flow) validateSuccess(step *Step) error {
	log := logger.WithField("method", "*Flow.validateSuccess")

	log.Debugf("called with step := [%s]", step)

	if len(step.OnSuccess) == 0 {
		return nil
	}
	for _, s := range step.OnSuccess {
		for _, value := range f.Steps {
			if value.Name == s.Name {
				return nil
			}
		}
	}
	return ErrMissingSuccessStep
}

// validateFailure checks if all required steps on failure
// exist
func (f *Flow) validateFailure(step *Step) error {
	log := logger.WithField("method", "*Flow.validateFailure")

	log.Debugf("called with step := [%s]", step)

	if len(step.OnFailure) == 0 {
		return nil
	}
	for _, s := range step.OnFailure {
		for _, value := range f.Steps {
			if value.Name == s.Name {
				return nil
			}
		}
	}
	return ErrMissingFailureStep
}
