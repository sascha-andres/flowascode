package flowascode

import (
	"strings"
	"testing"
)

func TestSuccessExist(t *testing.T) {
	yaml := `---
break_on_error: true
shell: bash

steps:
  - name: Default
    script:
      - echo 'hello'
    on_success:
      - name: test`

	flow, err := NewFromReader(strings.NewReader(yaml))
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if !flow.Steps[0].HasSuccess() {
		t.Log("expected success step, found none")
		t.Fail()
	}
}

func TestSuccessDoesNotExist(t *testing.T) {
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

	if flow.Steps[0].HasSuccess() {
		t.Log("expected no success step, found none")
		t.Fail()
	}
}
