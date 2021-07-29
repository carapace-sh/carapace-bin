package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

type TestOpts struct {
	Benchmark bool
	Example   bool
	Test      bool
}

func ActionTests(packages []string, opts TestOpts) carapace.Action {
	return carapace.ActionExecCommand("go", append([]string{"test", "-list=.*"}, packages...)...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if (opts.Benchmark && strings.HasPrefix(line, "Benchmark")) ||
				(opts.Example && strings.HasPrefix(line, "Example")) ||
				(opts.Test && strings.HasPrefix(line, "Test")) {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
