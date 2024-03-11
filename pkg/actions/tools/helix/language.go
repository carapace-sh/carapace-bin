package helix

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLanguages completes languages
//
//	bash (✘ bash-language-server)
//	c (✔ clangd)
func ActionLanguages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		cmd := "helix"
		if _, err := execlog.LookPath("helix"); err != nil {
			cmd = "hx"
		}
		return carapace.ActionExecCommand(cmd, "--health")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<language>[^ ]+) +(?P<lsp>(✘ |✓ )?[^ ]+)`)

			vals := make([]string, 0)
			for _, line := range lines[8:] {
				if match := r.FindStringSubmatch(line); match != nil {
					if lsp := match[2]; strings.HasPrefix(lsp, "✘") {
						vals = append(vals, match[1], lsp, style.Red)
					} else if strings.HasPrefix(lsp, "✓") {
						vals = append(vals, match[1], lsp, style.Green)
					} else {
						vals = append(vals, match[1], lsp, style.Yellow)
					}
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
