package actions

import (
	"github.com/carapace-sh/carapace"
	"regexp"
)

func ActionSpeed() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		r := regexp.MustCompile(`^(?P<digits>\d+).*$`)
		if matches := r.FindStringSubmatch(c.Value); matches != nil {
			return carapace.ActionValues("KB", "MB", "GB").Prefix(matches[1])
		}
		return carapace.ActionValues()
	})
}
