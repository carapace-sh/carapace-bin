package d2

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type ThemeOpts struct {
	Dark  bool
	Light bool
}

func (o ThemeOpts) Default() ThemeOpts {
	o.Dark = true
	o.Light = true
	return o
}

// ActionThemes completes themes
//
//	0 (Neutral Default)
//	1 (Neutral Grey)
func ActionThemes(opts ThemeOpts) carapace.Action {
	return carapace.ActionExecCommand("d2", "themes")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^- (?P<name>[^:]+): (?P<description>.*)$`)

		dark := false
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "Dark:" {
				dark = true
				continue
			}
			if matches := r.FindStringSubmatch(line); matches != nil {
				switch {
				case dark && opts.Dark:
					vals = append(vals, matches[2], matches[1], style.Blue)
				case !dark && opts.Light:
					vals = append(vals, matches[2], matches[1], style.Default)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("themes")
}
