package pamac

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionPackageGroups completes package groups
//
//	archlinux-tools ([group])
//	arduino ([group])
func ActionPackageGroups() carapace.Action {
	return carapace.ActionExecCommand("pamac", "list", "--groups")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, line, "[group]")
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionPackageSearch completes installable packages
//
//	git ([extra] the fast distributed version control system)
//	git-absorb ([extra] git commit --fixup, but automatic)
func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Value) == 0 {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("pamac", "search", "--aur", c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^ ]+) +(?P<version>[^ ]+) +(?P<installed>\[Installed\])? +(?P<repo>[^ ]+)$`)

			current := ""
			packages := make(map[string][]string)
			for _, line := range lines[:len(lines)-1] {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					current = matches[1]
					packages[current] = []string{fmt.Sprintf("[%v]", matches[4])}
				} else {
					packages[current] = append(packages[current], strings.TrimSpace(line))
				}
			}

			vals := make([]string, 0)
			for key, value := range packages {
				vals = append(vals, key, strings.Join(value, " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionInstalledPackages completes installed packages
//
//	a52dec (0.8.0-2)
//	aalib (1.4rc5-17)
func ActionInstalledPackages(explicit bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		arg := "-i"
		if explicit {
			arg = "-e"
		}
		return carapace.ActionExecCommand("pamac", "list", arg)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], fields[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})

	})
}
