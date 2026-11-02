package gh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type gitignoreTemplatesQuery []string

// ActionGitignoreTemplates completes gitignore templates
//
//	Actionscript
//	Android
func ActionGitignoreTemplates(opts HostOpts) carapace.Action {
	args := []string{"api", "--preview", "mercy", "gitignore/templates"}
	if opts.Host != "" {
		args = append(args, "--hostname", opts.Host)
	}
	return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
		var queryResult gitignoreTemplatesQuery
		if err := json.Unmarshal(output, &queryResult); err != nil {
			return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
		}
		return carapace.ActionValues(queryResult...)
	})
}
