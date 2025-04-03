package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time one of the other commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timeCmd).Standalone()

	rootCmd.AddCommand(timeCmd)

	carapace.Gen(timeCmd).PositionalCompletion(carapace.ActionCommands(rootCmd))

	carapace.Gen(timeCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
