package flowascode

import (
	"strings"
	"testing"
)

func TestNoDefault(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: hello world`

	flow, err := NewFromReader(strings.NewReader(yaml))
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	step, err := flow.GetDefault()
	if err == nil {
		t.Log("expected error, got none")
		t.Fail()
	}
	if nil != step {
		t.Logf("expected nil, got %s", step)
		t.Fail()
	}
}

func TestDefault(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: default`

	flow, err := NewFromReader(strings.NewReader(yaml))
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	step, err := flow.GetDefault()
	if err != nil {
		t.Log("expected no error, got one")
		t.Fail()
	}
	if nil == step {
		t.Log("expected one, got none")
		t.Fail()
	}
}
