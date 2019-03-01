package flowascode

import "testing"

func TestNewFromYAML(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: hello world`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if flow.BreakOnError != true {
		t.Log("expected break on error to be true, received false")
		t.Fail()
	}

	if len(flow.Steps) != 1 {
		t.Logf("expected one step, got: %d", len(flow.Steps))
		t.Fail()
	}
}
