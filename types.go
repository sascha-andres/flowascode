package flowascode

type (
	// Flow represents a complete flow
	Flow struct {
		BreakOnError bool   `yaml:"break_on_error"` // BreakOnError breaks if a script returns a non zero result
		Steps        []Step `yaml:"steps"`          // Steps is the collection of steps in the flow
	}

	// Step is a single execution point in a flow
	Step struct {
		Name      string       `yaml:"name"` // Name of a step, also used to reference from C(Descendant)
		Script    []string     // Script to execute
		OnSuccess []Descendant // OnSuccess is a list of steps to execute if the script returns with 0
		OnFailure []Descendant // OnFailure is a list of steps to execute if the scripts returns not a 0
	}

	// Descendant represents on step after the current step
	// that is it is a reference to a number of other steps
	Descendant struct {
		Name      string            // Name is the referenced step
		Variables map[string]string // Variables is a list of variables that are set as environment variables
	}
)
