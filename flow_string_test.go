package flowascode

import (
	"fmt"
	"testing"
)

func TestFlow_String(t *testing.T) {
	yaml := `---
break_on_error: true
name: test
description: test

steps:
  - name: Default`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if "test\ntest" != fmt.Sprintf("%s", flow) {
		t.Logf("expected [Default] got [%s]", flow)
		t.Fail()
	}
}
