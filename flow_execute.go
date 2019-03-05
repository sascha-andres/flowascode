package flowascode

import "fmt"

// Execute takes a step name and starts to execute from there
// if no ma,e is given it tries to get the default namespace
func (f *Flow) Execute(name string) error {
	var (
		step *Step
		err  error
	)
	if name == "" && f.HasDefault() {
		step, err = f.GetDefault()
		if err != nil {
			return nil
		}
	}
	fmt.Println(step)
	return Error("not yet implemented")
}
