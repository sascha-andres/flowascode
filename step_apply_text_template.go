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
	"bytes"
	"text/template"
)

func (s *Step) applyTextTemplate(f *Flow, value string) string {
	log := logger.WithField("method", "*Step.applyTextTemplate")

	log.Debugf("called with s := [%s]", s)
	log.Debugf("called with f := [%s]", f)
	log.Debugf("called with value := [%s]", value)

	t := template.Must(template.New("").Parse(value))
	var tpl bytes.Buffer
	_ = t.Execute(&tpl, struct {
		Flow *Flow
		Step *Step
	}{
		Flow: f,
		Step: s,
	})
	return tpl.String()
}
