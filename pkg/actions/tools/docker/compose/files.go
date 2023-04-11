package compose

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type ServicePathOpts struct {
	Files   []string
	Service string
	Index   int
}

// ActionFiles completes files within a service container
//
//	/home
//	/etc
func ActionFiles(opts ServicePathOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if opts.Index == 0 {
			opts.Index = 1 // default to 1
		}

		path := filepath.Dir(c.Value)
		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return actionExecCompose(opts.Files, "exec", "--no-TTY", "--index", strconv.Itoa(opts.Index), opts.Service, "ls", `-1`, `-p`, path)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}).NoSpace().StyleF(style.ForPathExt)
	})
}
