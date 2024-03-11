package supervisor

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionProcesses completes processes
//
//	foo (pid 19055, uptime 1:03:50)
//	sleep:0 (Mar 10 01:01 PM)
func ActionProcesses(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			var err error
			if path, err = configPath(c); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		return carapace.ActionExecCommandE("supervisorctl", "--configuration", path, "status")(func(output []byte, err error) carapace.Action {
			if err != nil {
				const NOT_RUNNING = 3
				if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != NOT_RUNNING {
					if len(exitErr.Stderr) > 0 {
						return carapace.ActionMessage(string(exitErr.Stderr))
					}
					if len(output) > 0 {
						return carapace.ActionMessage(string(output))
					}
					return carapace.ActionMessage(string(output))
				}
			}

			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			r := regexp.MustCompile(`^(?P<name>[^ ]+) +(?P<status>[^ ]+) +(?P<description>.*)$`)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil {
					vals = append(vals, matches[1], matches[3], style.ForKeyword(matches[2], c))
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
