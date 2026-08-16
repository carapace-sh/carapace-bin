package devenv

import (
	"os"
	"regexp"

	"github.com/carapace-sh/carapace"
)

var containerPattern = regexp.MustCompile(`(?m)^\s*containers\.([\w-]+)`)

// ActionContainers completes containers.
// Custom containers are scraped from devenv.nix as evaluating them is only possible on linux.
//
//	processes (start the processes)
//	shell (enter the developer environment)
func ActionContainers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// devenv always defines these two
		descriptions := map[string]string{
			"processes": "start the processes",
			"shell":     "enter the developer environment",
		}

		if file, err := configFile(c); err == nil {
			content, err := os.ReadFile(file)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			for _, match := range containerPattern.FindAllStringSubmatch(string(content), -1) {
				if _, ok := descriptions[match[1]]; !ok {
					descriptions[match[1]] = ""
				}
			}
		}

		vals := make([]string, 0)
		for name, description := range descriptions {
			vals = append(vals, name, description)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("containers")
}
