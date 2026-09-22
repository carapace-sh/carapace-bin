package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Runs the script defined by the current workspace's Anchor.toml",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		anchor.ActionScripts(),
	)
}
