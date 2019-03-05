package flowascode

import (
	"github.com/hashicorp/go-multierror"
	"testing"
)

func TestNoSuchStep(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: Default`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if err := flow.ValidateStep("test"); err == nil {
		t.Log("no step [test] but step is valid")
		t.Fail()
	} else {
		if err != ErrMissingStep {
			t.Logf("expected [%s] got [%s]", ErrMissingStep, err)
			t.Fail()
		}
	}
}

func TestMissingSuccessStep(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: missing-success
    on_success:
      - name: next
    script:
      - echo a`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if err := flow.ValidateStep("missing-success"); err == nil {
		t.Log("[missing-success] has missing dependency but sent as valid")
		t.Fail()
	} else {
		if merr, ok := err.(*multierror.Error); ok {
			if merr.Len() != 1 {
				t.Logf("expected [%s] got [%s]", ErrMissingSuccessStep, err)
				t.Fail()
			} else {
				if merr.Errors[0] != ErrMissingSuccessStep {
					t.Logf("expected [%s] got [%s]", ErrMissingSuccessStep, err)
					t.Fail()
				}
			}
		}
	}
}

func TestSuccessStep(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: success
    on_success:
      - name: next
    script:
      - echo a
  - name: next`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if err := flow.ValidateStep("success"); err != nil {
		t.Log("[success] should be valid but listed as invalid")
		t.Logf("error is: %s", err)
		t.Fail()
	}
}

func TestMissingFailureStep(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: missing-failure
    on_failure:
      - name: next
    script:
      - echo a`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if err := flow.ValidateStep("missing-failure"); err == nil {
		t.Log("[missing-failure] has missing dependency but sent as valid")
		t.Fail()
	} else {
		if merr, ok := err.(*multierror.Error); ok {
			if merr.Len() != 1 {
				t.Logf("expected [%s] got [%s]", ErrMissingFailureStep, err)
				t.Fail()
			} else {
				if merr.Errors[0] != ErrMissingFailureStep {
					t.Logf("expected [%s] got [%s]", ErrMissingFailureStep, err)
					t.Fail()
				}
			}
		}
	}
}

func TestFailureStep(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: failure
    on_failure:
      - name: next
    script:
      - echo a
  - name: next`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if err := flow.ValidateStep("failure"); err != nil {
		t.Log("[failure] should be valid but listed as invalid")
		t.Logf("error is: %s", err)
		t.Fail()
	}
}

func TestMissingScript(t *testing.T) {
	yaml := `---
break_on_error: true

steps:
  - name: missing-script`

	flow, err := NewFromYAML(yaml)
	if err != nil {
		t.Logf("expected no error got: %s", err)
		t.Fail()
	}

	if err := flow.ValidateStep("missing-script"); err == nil {
		t.Log("[missing-script] has missing script but sent as valid")
		t.Fail()
	} else {
		if merr, ok := err.(*multierror.Error); ok {
			if merr.Len() != 1 {
				t.Logf("expected [%s] got [%s]", ErrNoScript, err)
				t.Fail()
			} else {
				if merr.Errors[0] != ErrNoScript {
					t.Logf("expected [%s] got [%s]", ErrNoScript, err)
					t.Fail()
				}
			}
		}
	}
}
