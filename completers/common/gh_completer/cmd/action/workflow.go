package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type WorkflowOpts struct {
	Enabled  bool
	Disabled bool
	Id       bool
	Name     bool
}

// ActionWorkflows completes workflows
func ActionWorkflows(cmd *cobra.Command, opts WorkflowOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionWorkflows(gh.WorkflowOpts{
		Host:     repo.Host,
		Owner:    repo.Owner,
		Name:     repo.Name,
		Enabled:  opts.Enabled,
		Disabled: opts.Disabled,
		ById:     opts.Id,
		ByName:   opts.Name,
	})
}
