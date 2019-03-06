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

package main

import (
	"github.com/google/gops/agent"
	"github.com/integrii/flaggy"
	"github.com/sascha-andres/flowascode"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	flow   string
	name   string
	debug  bool
	logger *logrus.Entry
)

func main() {
	log := logger.WithField("method", "main")

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	if err := validate(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	f, err := os.Open(flow)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()

	flow, err := flowascode.NewFromReader(f)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if debug {
		flowascode.SetLogLevel(logrus.DebugLevel)
	}

	err = flow.Execute(name, nil)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func validate() error {
	log := logger.WithField("method", "validate")
	log.Debug("called")

	if "" == strings.Trim(name, " ") {
		return flowascode.Error("no step name provided")
	}
	if "" == strings.Trim(flow, " ") {
		return flowascode.Error("no flow file name provided")
	}
	return nil
}

// init is called as soon as something in the package is used for the first time
// used here to get command line arguments
func init() {
	logger = logrus.WithField("package", "main")

	flaggy.Bool(&debug, "d", "debug", "turn on debug logging")
	flaggy.String(&flow, "f", "flow", "provide path to flow definition")
	flaggy.String(&name, "n", "name", "name of step to execute")

	flaggy.Parse()
}
