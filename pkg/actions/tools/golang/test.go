package golang

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type TestOpts struct {
	Packages  []string
	Benchmark bool
	Example   bool
	Test      bool
}

func (o TestOpts) Default() TestOpts {
	o.Benchmark = true
	o.Example = true
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
			} else if opts.Test && strings.HasPrefix(line, "Test") {
				vals = append(vals, line, style.Default)
			}
		}
		return carapace.ActionStyledValues(vals...)
	})
}
