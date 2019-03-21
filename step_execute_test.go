package flowascode

import "testing"

func TestNilScript(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: Default`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	err = flow.Steps[0].Execute("", nil)
	if err == nil {
		t.Log("expected error, got none")
		t.Fail()
	}

	if err != ErrNoScript {
		t.Logf("expected [%s], got [%s]", ErrNoScript, err)
		t.Fail()
	}
}

func TestEmptyScript(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: Default
    script: []`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	err = flow.Steps[0].Execute("", nil)
	if err == nil {
		t.Log("expected error, got none")
		t.Fail()
	}

	if err != ErrNoScript {
		t.Logf("expected [%s], got [%s]", ErrNoScript, err)
		t.Fail()
	}
}

func TestScript(t *testing.T) {
	yaml := `---
break_on_error: true
shell: bash

steps:
  - name: Default
    script:
      - echo 'hello'`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	err = flow.Steps[0].Execute(flow.Shell, nil)
	if err != nil {
		t.Logf("expected no error, got [%s]", err)
		t.Fail()
	}
}
