package actions

import "github.com/carapace-sh/carapace"

func ActionSpeed() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Value) > 0 && c.Value[0] >= '0' && c.Value[0] <= '9' {
			num := ""
			for _, char := range c.Value {
				if char >= '0' && char <= '9' {
					num += string(char)
				} else {
					break
				}
			}
			return carapace.ActionValues("KB", "MB", "GB").Invoke(c).Prefix(num).ToA()
		}
		return carapace.ActionValues()
	})
}
