package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

func ActionIssues(cmd *cobra.Command, open bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := tea.IssueOpts{}
		if f := cmd.Flag("login"); f != nil {
			opts.Login = f.Value.String()
		}
		if f := cmd.Flag("remote"); f != nil {
			opts.Remote = f.Value.String()
		}
		if f := cmd.Flag("repo"); f != nil {
			opts.Repo = f.Value.String()
		}
		opts.Open = open
		opts.Closed = !open
		return tea.ActionIssues(opts)
	})
}

func ActionPullrequests(cmd *cobra.Command, open bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := tea.PullrequestOpts{}
		if f := cmd.Flag("login"); f != nil {
			opts.Login = f.Value.String()
		}
		if f := cmd.Flag("remote"); f != nil {
			opts.Remote = f.Value.String()
		}
		if f := cmd.Flag("repo"); f != nil {
			opts.Repo = f.Value.String()
		}
		opts.Open = open
		opts.Closed = !open
		return tea.ActionPullrequests(opts)
	})
}

func ActionLabels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := tea.RepoOpts{}
		if f := cmd.Flag("login"); f != nil {
			opts.Login = f.Value.String()
		}
		if f := cmd.Flag("remote"); f != nil {
			opts.Remote = f.Value.String()
		}
		if f := cmd.Flag("repo"); f != nil {
			opts.Repo = f.Value.String()
		}
		return tea.ActionLabels(opts)
	})
}
