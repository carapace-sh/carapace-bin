package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eCmd = &cobra.Command{
	Use:   "e",
	Short: "Extract files from archive (without using directory names)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eCmd).Standalone()

	addCommonFlags(eCmd)
	addCommonFlagCompletions(eCmd)
	rootCmd.AddCommand(eCmd)

	carapace.Gen(eCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(eCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
