package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionGitignoreTemplates completes gitignore templates
func ActionGitignoreTemplates(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionGitignoreTemplates(gh.HostOpts{Host: repo.Host})
}
