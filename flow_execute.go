package flowascode

import "github.com/sirupsen/logrus"

// Execute takes a step name and starts to execute from there
// if no ma,e is given it tries to get the default namespace
func (f *Flow) Execute(name string) error {
	log := logrus.
		WithField("package", "flowascode").
		WithField("method", "Execute")

	log.Debugf("called with name := [%s]", name)

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
		log.Errorf("error finding step: %s", err)
		return err
	}
	err = f.ValidateStep(name)
	if err != nil {
		log.Errorf("step not valid: %s", err)
		return err
	}
	err = step.Execute(f.Shell)
	if err != nil {
		log.Errorf("step execution failed: %s", err)
		for _, value := range step.OnFailure {
			_ = f.Execute(value.Name)
		}
	} else {
		log.Info("step execution succeeded")
		for _, value := range step.OnSuccess {
			_ = f.Execute(value.Name)
		}
	}
	return err
}
