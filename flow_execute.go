package flowascode

// Execute takes a step name and starts to execute from there
// if no ma,e is given it tries to get the default namespace
func (f *Flow) Execute(name string) error {
	var (
		step *Step
		err  error
	)
	if name == "" && f.HasDefault() {
		step, err = f.GetDefault()
	} else {
		step, err = f.GetStep(name)
	}
	if err != nil {
		return err
	}
	err = f.ValidateStep(name)
	if err != nil {
		return err
	}
	err = step.Execute(f.Shell)
	if err != nil {
		for _, value := range step.OnFailure {
			err = f.Execute(value.Name)
		}
	} else {
		for _, value := range step.OnSuccess {
			err = f.Execute(value.Name)
		}
	}
	return err
}
