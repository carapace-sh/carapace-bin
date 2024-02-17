package mvn

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache/key"
	"github.com/rsteube/carapace/pkg/util"
)

// ActionProjects completes projects
//
//	core
//	another
func ActionProjects(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"-Dexec.executable=echo", "-Dexec.args=${project.artifactId}", "exec:exec", "-q"}
		if file != "" {
			args = append(args, "--file", file)
		} else {
			var err error
			if file, err = util.FindReverse(c.Dir, "pom.xml"); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}
		return carapace.ActionExecCommand("mvn", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}).Cache(-1, key.FileStats(file))
	})
}
