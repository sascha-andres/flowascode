package flowascode

import (
	"strings"
	"testing"
)

func TestFailureExist(t *testing.T) {
	yaml := `---
break_on_error: true
shell: bash

steps:
  - name: Default
    script:
      - echo 'hello'
    on_failure:
      - name: test`

	flow, err := NewFromReader(strings.NewReader(yaml))
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if !flow.Steps[0].HasFailure() {
		t.Log("expected success step, found none")
		t.Fail()
	}
}

func TestFailureDoesNotExist(t *testing.T) {
	yaml := `---
break_on_error: true
shell: bash

steps:
  - name: Default
    script:
      - echo 'hello'`

	flow, err := NewFromReader(strings.NewReader(yaml))
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if flow.Steps[0].HasFailure() {
		t.Log("expected no failure step, found none")
		t.Fail()
	}
}
