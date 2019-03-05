package flowascode

import "testing"

func TestFlow_HasDefault_UpperCase(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: Default`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if !flow.HasDefault() {
		t.Log("default step not detected")
		t.Fail()
	}
}

func TestFlow_HasDefault_LowerCase(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: default
  - name: hello`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if !flow.HasDefault() {
		t.Log("default step not detected")
		t.Fail()
	}
}

func TestFlow_Not_HasDefault(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - Name: test`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if flow.HasDefault() {
		t.Log("default step detected without expecting")
		t.Fail()
	}
}
