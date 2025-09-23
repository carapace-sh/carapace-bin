package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var agentTask_viewCmd = &cobra.Command{
	Use:   "view [<session-id> | <pr-number> | <pr-url> | <pr-branch>]",
	Short: "View an agent task session (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentTask_viewCmd).Standalone()

	agentTask_viewCmd.Flags().Bool("follow", false, "Follow agent session logs")
	agentTask_viewCmd.Flags().Bool("log", false, "Show agent session logs")
	agentTask_viewCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	agentTask_viewCmd.Flags().BoolP("web", "w", false, "Open agent task in the browser")
	agentTaskCmd.AddCommand(agentTask_viewCmd)

	carapace.Gen(agentTask_viewCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionOwnerRepositories(gh.HostOpts{}),
	})

	// TODO positional completion
}
