package flowascode

import (
	"os"
	"strings"
	"testing"
)

func TestNewFromReader(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - Name: hello world`

	flow, err := NewFromReader(strings.NewReader(yaml))
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

func TestNewFromReaderNilReader(t *testing.T) {
	_, err := NewFromReader(nil)
	if err == nil {
		t.Log("expected error got none")
		t.Fail()
	}
}

func TestNewFromReaderErrorReader(t *testing.T) {
	file, _ := os.Open("jkhgkjhkjh")
	_, err := NewFromReader(file)
	if err == nil {
		t.Log("expected error got none")
		t.Fail()
	}
}
