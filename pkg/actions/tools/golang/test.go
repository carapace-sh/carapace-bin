package golang

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type TestOpts struct {
	Packages  []string
	Benchmark bool
	Example   bool
	Fuzz      bool
	Test      bool
}

func (o TestOpts) Default() TestOpts {
	o.Benchmark = true
	o.Example = true
	o.Fuzz = true
	o.Test = true
	return o
}

// ActionTests completes tests
//
//	TestActionFiles
//	TestActionFilesChdir
func ActionTests(opts TestOpts) carapace.Action {
	return carapace.ActionExecCommand("go", append([]string{"test", "-list=.*"}, opts.Packages...)...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if opts.Benchmark && strings.HasPrefix(line, "Benchmark") {
				vals = append(vals, line, style.Blue)
			} else if opts.Example && strings.HasPrefix(line, "Example") {
				vals = append(vals, line, style.Green)
			} else if opts.Fuzz && strings.HasPrefix(line, "Fuzz") {
				vals = append(vals, line, style.Magenta)
			} else if opts.Test && strings.HasPrefix(line, "Test") {
				vals = append(vals, line, style.Default)
			}
		}
		return carapace.ActionStyledValues(vals...)
	})
}
