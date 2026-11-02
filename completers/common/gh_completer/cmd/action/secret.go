package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type SecretOpts struct {
	App string
	Org string
	Env string
}

// ActionSecrets completes secrets
func ActionSecrets(cmd *cobra.Command, opts SecretOpts) carapace.Action {
	repo := ""
	if flag := cmd.Flag("repo"); flag != nil {
		repo = flag.Value.String()
	}
	return gh.ActionSecrets(gh.SecretsOpts{
		Repo: repo,
		App:  opts.App,
		Org:  opts.Org,
		Env:  opts.Env,
	})
}
