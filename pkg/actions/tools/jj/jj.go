package jj

import "github.com/carapace-sh/carapace"

func actionExecJJ(arg ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionExecCommand("jj", append([]string{"--color", "never"}, arg...)...)(func(output []byte) carapace.Action {
			return f(output)
		})
	}
}
