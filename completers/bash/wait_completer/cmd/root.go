package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for job completion and return exit status",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Job-Control-Builtins.html#index-wait",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("f", "f", false, "wait for the specified ID to terminate instead of waiting for it to change status")
	rootCmd.Flags().BoolS("n", "n", false, "wait for a single job from the list of IDs to complete")
	rootCmd.Flags().StringS("p", "p", "", "assign the process or job identifier of the job to variable var")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.Batch(
			shell.ActionJobSpecs(),
			ps.ActionProcessIds(),
		).ToA().FilterArgs(),
	)
}
