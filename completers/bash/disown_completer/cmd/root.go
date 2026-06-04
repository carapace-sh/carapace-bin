package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "disown",
	Short: "Remove jobs from current shell",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Job-Control-Builtins.html#index-disown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "remove all jobs if jobspec is not supplied")
	rootCmd.Flags().BoolS("h", "h", false, "mark each jobspec so that SIGHUP is not sent to the job if the shell receives a SIGHUP")
	rootCmd.Flags().BoolS("r", "r", false, "remove only running jobs")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.Batch(
			shell.ActionJobSpecs(),
			ps.ActionProcessIds(),
		).ToA().FilterArgs(),
	)
}
