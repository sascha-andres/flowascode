package flowascode

import "fmt"

// String method is used to print values passed as an operand
func (f Flow) String() string {
	return fmt.Sprintf("%s\n%s", f.Name, f.Description)
}
