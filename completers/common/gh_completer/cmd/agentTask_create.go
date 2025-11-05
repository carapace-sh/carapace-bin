package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var agentTask_createCmd = &cobra.Command{
	Use:   "create [<task description>] [flags]",
	Short: "Create an agent task (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentTask_createCmd).Standalone()

	agentTask_createCmd.Flags().StringP("base", "b", "", "Base branch for the pull request (use default branch if not provided)")
	agentTask_createCmd.Flags().StringP("custom-agent", "a", "", "Use a custom agent for the task. e.g., use 'my-agent' for the 'my-agent.md' agent")
	agentTask_createCmd.Flags().Bool("follow", false, "Follow agent session logs")
	agentTask_createCmd.Flags().StringP("from-file", "F", "", "Read task description from `file` (use \"-\" to read from standard input)")
	agentTask_createCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	agentTaskCmd.AddCommand(agentTask_createCmd)

	carapace.Gen(agentTask_createCmd).FlagCompletion(carapace.ActionMap{
		"from-file": carapace.ActionFiles(),
		"repo":      gh.ActionOwnerRepositories(gh.HostOpts{}),
	})
}
