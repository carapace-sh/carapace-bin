package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Display status of jobs",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Job-Control-Builtins.html#index-jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("l", "l", false, "list process IDs in addition to the normal information")
	rootCmd.Flags().BoolS("n", "n", false, "display only jobs that have changed status since the user was last notified of their status")
	rootCmd.Flags().BoolS("p", "p", false, "display only the process ID of the job's process group leader")
	rootCmd.Flags().BoolS("r", "r", false, "display only running jobs")
	rootCmd.Flags().BoolS("s", "s", false, "display only stopped jobs")
	rootCmd.Flags().BoolS("x", "x", false, "replace job spec with process group ID in command")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionJobSpecs().FilterArgs(),
	)
}
