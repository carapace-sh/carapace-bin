package ssh

import (
	"bufio"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/internal/actions/net/ssh"
)

// ActionHosts completes ssh hosts
func ActionHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := c.Abs("~/.ssh/config")
		if err != nil {
			return carapace.ActionValues()
		}

		file, err := os.Open(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		defer file.Close()

		vals := make([]string, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if line := scanner.Text(); strings.HasPrefix(line, "Host ") {
				for _, field := range strings.Fields(strings.TrimPrefix(line, "Host ")) {
					if !strings.ContainsAny(field, `\*{}[]+`) {
						vals = append(vals, field)
					}
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("ssh hosts")
}

func init() {
	ssh.ActionHosts = ActionHosts
}
