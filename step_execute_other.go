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

// +build !windows

package flowascode

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// executeOS implementation for Linux
func (s *Step) executeOS(pathToShell string, variables []string) error {
	log := logger.WithField("method", "*Step.executeOS[other]")

	log.Debugf("called with pathToShell := [%s]", pathToShell)

	var (
		shellCommand string
		err          error
	)

	// check if shell exist
	if shellCommand, err = exec.LookPath(pathToShell); err != nil {
		return err
	}

	temporaryFile, def, err := createTempFile()
	if def != nil {
		defer def()
	}
	if err != nil {
		return err
	}

	if err := s.fillFile(temporaryFile); err != nil {
		return err
	}

	return executeScript(shellCommand, temporaryFile.Name(), variables)
}

// createTempFile creates a temporary file in the temp directory
func createTempFile() (*os.File, func(), error) {
	dir, err := ioutil.TempDir("", "flowascode")
	if err != nil {
		return nil, nil, err
	}

	temporaryFile, err := ioutil.TempFile(dir, "example")
	if err != nil {
		return nil, nil, err
	}

	return temporaryFile, func() {
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Println(err)
		}
	}, nil
}

// fillFile writes the data to a file
func (s *Step) fillFile(temporaryFile *os.File) error {
	script := strings.Join(s.Script, "\n")
	if _, err := temporaryFile.Write([]byte(script)); err != nil {
		return err
	}
	err := temporaryFile.Close()
	return err
}

// executeScript runs the created script
func executeScript(shellCommand, temporaryFile string, variables []string) error {
	cmd := exec.Command(shellCommand, temporaryFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if 0 != len(variables) {
		cmd.Env = variables
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	_ = cmd.Wait()
	if cmd.ProcessState.ExitCode() == 0 {
		return nil
	}
	return errors.New(fmt.Sprintf("exit status: %d", cmd.ProcessState.ExitCode()))
}
