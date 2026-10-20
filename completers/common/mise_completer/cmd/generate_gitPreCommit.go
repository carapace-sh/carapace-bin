package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var generate_gitPreCommitCmd = &cobra.Command{
	Use:     "git-pre-commit",
	Short:   "Generate a git hook that runs a mise task",
	Aliases: []string{"pre-commit"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_gitPreCommitCmd).Standalone()

	generate_gitPreCommitCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_gitPreCommitCmd.Flags().String("hook", "pre-commit", "Which git hook to write, such as pre-push")
	generate_gitPreCommitCmd.Flags().StringP("task", "t", "pre-commit", "The task to run when the hook is triggered")
	generate_gitPreCommitCmd.Flags().BoolP("write", "w", false, "Write the hook to the hooks directory and make it executable")
	generateCmd.AddCommand(generate_gitPreCommitCmd)

	carapace.Gen(generate_gitPreCommitCmd).FlagCompletion(carapace.ActionMap{
		"task": mise.ActionTasks(),
	})
}
