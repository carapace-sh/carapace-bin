package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var generate_githubActionCmd = &cobra.Command{
	Use:   "github-action",
	Short: "Generate a GitHub Actions workflow that runs a mise task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_githubActionCmd).Standalone()

	generate_githubActionCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_githubActionCmd.Flags().String("name", "ci", "Name of the workflow and its file")
	generate_githubActionCmd.Flags().StringP("task", "t", "ci", "The task to run when the workflow is triggered")
	generate_githubActionCmd.Flags().BoolP("write", "w", false, "Write the workflow to `.github/workflows/<NAME>.yml` instead of printing it")
	generateCmd.AddCommand(generate_githubActionCmd)

	carapace.Gen(generate_githubActionCmd).FlagCompletion(carapace.ActionMap{
		"task": mise.ActionTasks(),
	})
}
