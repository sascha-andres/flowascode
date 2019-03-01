package flowascode

import (
	"fmt"
	"testing"
)

func TestStep_String(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: Default`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if len(flow.Steps) != 1 {
		t.Logf("expected 1 step, got %d", len(flow.Steps))
		t.Fail()
	}

	if "Default" != fmt.Sprintf("%s", flow.Steps[0]) {
		t.Logf("expected [Default] got [%s]", flow.Steps[0])
		t.Fail()
	}
}
