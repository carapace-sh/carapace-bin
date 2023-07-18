package styles

import "github.com/rsteube/carapace/pkg/style"

var Golang = struct {
	Constant string `description:"go constant"`
	Field    string `description:"go field"`
	Function string `description:"go function"`
	Type     string `description:"go type"`
	Variable string `description:"go variable"`
}{
	Constant: style.Magenta,
	Field:    style.Green,
	Function: style.Blue,
	Type:     style.Yellow,
	Variable: style.Default,
}

func init() {
	style.Register("golang", &Golang)
}
