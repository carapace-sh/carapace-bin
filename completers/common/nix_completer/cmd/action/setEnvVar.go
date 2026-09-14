package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
)

// ActionSetEnvVar completes name and value for the --set-env-var flag
//
//	HOME /home/user
//	PATH /usr/bin:/bin
func ActionSetEnvVar() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				env.ActionNames(),
				os.ActionEnvironmentVariables(),
			).ToA()
		default:
			return env.ActionValues(c.Parts[0])
		}
	})
}
