package hl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionThemes completes color themes
//
//	ayu-dark-24
//	classic
//	uni
func ActionThemes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("hl", "--list-themes")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
