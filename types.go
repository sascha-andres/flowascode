package flowascode

import "github.com/sirupsen/logrus"

type (

	// Error type for constant error definitions
	Error string

	// Flow represents a complete flow
	Flow struct {
		BreakOnError bool   `yaml:"break_on_error"` // BreakOnError breaks if a script returns a non zero result
		Name         string `yaml:"name"`           // Name is a name for the flow
		Description  string `yaml:"description"`    // Description is a more descriptive text what the flow does
		Steps        []Step `yaml:"steps"`          // Steps is the collection of steps in the flow
		Shell        string `yaml:"shell"`          // Shell is the path or binary name for the shell to use to execute the script
	}

	// Step is a single execution point in a flow
	Step struct {
		Name      string       `yaml:"name"`       // Name of a step, also used to reference from C(Descendant)
		Script    []string     `yaml:"script"`     // Script to execute
		OnSuccess []Descendant `yaml:"on_success"` // OnSuccess is a list of steps to execute if the script returns with 0
		OnFailure []Descendant `yaml:"on_failure"` // OnFailure is a list of steps to execute if the scripts returns not a 0
	}

	// Descendant represents on step after the current step
	// that is it is a reference to a number of other steps
	Descendant struct {
		Name      string            `yaml:"name"`      // Name is the referenced step
		Variables map[string]string `yaml:"variables"` // Variables is a list of variables that are set as environment variables
	}
)

var (
	logger *logrus.Entry
)

func (e Error) Error() string { return string(e) }

func init() {
	logger = logrus.WithField("package", "flowascode")
}
