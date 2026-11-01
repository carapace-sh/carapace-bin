package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Walk commit log and show metadata for each commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().BoolP("help", "h", false, "Print help")
	logCmd.Flags().Bool("mo", false, "Only show commits that have metadata")
	logCmd.Flags().StringS("n", "n", "20", "Number of commits to show (default: 20)")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
