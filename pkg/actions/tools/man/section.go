package man

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSections completes sections of given manpage (or all if empty).
//
//	1 (General commands)
//	2 (System calls)
func ActionSections(manpage string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		a := carapace.ActionValuesDescribed(
			"1", "General commands",
			"2", "System calls",
			"3", "Library functions",
			"4", "Special files",
			"5", "File formats and conventions",
			"6", "Games and screensavers",
			"7", "Miscellaneous",
			"8", "System administration commands and daemons",
		)

		if manpage != "" {
			output, err := c.Command("man", "-f", manpage).Output()
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			r := regexp.MustCompile(`^` + regexp.QuoteMeta(manpage) + ` \((?P<section>[^]]+)\).*$`)
			sections := make([]string, 0)
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil {
					sections = append(sections, matches[1][:1]) // only use first character
				}
			}
			a = a.Retain(sections...)
		}
		return a
	})
}
