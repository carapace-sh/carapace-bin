package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Opens the file in playground, an online web viewer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(playCmd).Standalone()

	rootCmd.AddCommand(playCmd)

	carapace.Gen(playCmd).PositionalCompletion(
		carapace.ActionFiles(".d2"),
	)
}
