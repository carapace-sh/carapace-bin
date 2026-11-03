package gh

import (
	"github.com/carapace-sh/carapace"
	"strings"
)

// ActionSecretFields completes secret fields
//
//	name
//	visibility
func ActionSecretFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"numSelectedRepos",
		"selectedReposURL",
		"updatedAt",
		"visibility",
	)
}

type SecretOpts struct {
	// Repo in `[HOST/]OWNER/REPO` format (the gh CLI passes it as-is)
	Repo string

	// App filters secrets to the given app (actions, codespaces, dependabot)
	App string
	// Org restricts to organization secrets
	Org string
	// Env restricts to environment secrets (requires Repo)
	Env string
}

func (o SecretOpts) Default() SecretOpts {
	return o
}

func (o SecretOpts) format() []string {
	args := make([]string, 0)
	if o.Repo != "" {
		args = append(args, "--repo", o.Repo)
	}
	if o.App != "" {
		args = append(args, "--app", o.App)
	}
	if o.Org != "" {
		args = append(args, "--org", o.Org)
	}
	if o.Env != "" {
		args = append(args, "--env", o.Env)
	}
	return args
}

// ActionSecrets completes secrets
//
//	GH_TOKEN (gh_pat_1234...)
func ActionSecrets(opts SecretOpts) carapace.Action {
	return carapace.ActionExecCommand("gh", append([]string{"secret", "list"}, opts.format()...)...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0, len(lines)-1)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, "\t", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
