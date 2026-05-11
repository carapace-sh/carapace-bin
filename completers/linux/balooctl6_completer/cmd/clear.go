package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Forget the specified files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCmd).Standalone()

	rootCmd.AddCommand(clearCmd)

	carapace.Gen(clearCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
