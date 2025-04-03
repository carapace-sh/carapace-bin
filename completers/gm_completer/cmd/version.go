package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "obtain release version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).PositionalCompletion(carapace.ActionCommands(rootCmd))
}
