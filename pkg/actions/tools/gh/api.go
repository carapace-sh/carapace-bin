package gh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

func apiV3Action(opts RepoOpts, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{
			"api",
			"--preview", "mercy",
			query,
		}
		if opts.Host != "" {
			args = append(args, "--hostname", opts.Host)
		}

		return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}
