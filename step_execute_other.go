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
func (s *Step) executeOS(pathToShell string) error {
	log := logger.WithField("method", "executeOS[other]")

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

	return executeScript(shellCommand, temporaryFile.Name())
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
func executeScript(shellCommand, temporaryFile string) error {
	cmd := exec.Command(shellCommand, temporaryFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
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
