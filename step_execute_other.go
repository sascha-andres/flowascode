// +build !windows

package flowascode

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// executeOS implementation for Linux
func (s *Step) executeOS(pathToShell string) error {
	var (
		shellCommand string
		err          error
	)

	// check if shell exist
	if shellCommand, err = exec.LookPath(pathToShell); err != nil {
		return err
	}

	dir, err := ioutil.TempDir("", "flowascode")
	if err != nil {
		return err
	}
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Println(err)
		}

	}()

	temporaryFile, err := ioutil.TempFile(dir, "example")
	if err != nil {
		log.Fatal(err)
	}

	script := strings.Join(s.Script, "\n")
	if _, err := temporaryFile.Write([]byte(script)); err != nil {
		return err
	}
	err = temporaryFile.Close()
	if err != nil {
		return err
	}
	cmd := exec.Command(shellCommand, temporaryFile.Name())
	err = cmd.Start()
	if err != nil {
		return err
	}
	_ = cmd.Wait()
	if cmd.ProcessState.ExitCode() == 0 {
		return nil
	}
	return errors.New(fmt.Sprintf("exit status: %d", cmd.ProcessState.ExitCode()))
}
