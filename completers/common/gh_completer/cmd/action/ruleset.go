package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionRulesets completes rulesets
func ActionRulesets(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionRulesets(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
