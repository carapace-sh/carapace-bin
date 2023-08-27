package bat

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionLanguages completes languages
//
//	Batch File (bat,cmd)
//	BibTeX (bib)
func ActionLanguages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("bat", "--list-languages")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			values := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				splitted := strings.Split(line, ":")
				if len(splitted) == 1 {
					values = append(values, splitted[0], "", style.Default)
				} else if len(splitted) > 1 {
					values = append(values, splitted[0], splitted[1], style.ForPathExt("."+strings.SplitN(splitted[1], ",", 2)[0], c))
				}
			}
			return carapace.ActionStyledValuesDescribed(values...)
		})
	})
}
