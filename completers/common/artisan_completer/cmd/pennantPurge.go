package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pennantPurgeCmd = &cobra.Command{
	Use:     "pennant:purge [--except [EXCEPT]] [--except-registered] [--store [STORE]] [--] [<features>...]",
	Short:   "Delete Pennant features from storage",
	Aliases: []string{"pennant:clear"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pennantPurgeCmd).Standalone()

	pennantPurgeCmd.Flags().String("except", "", "The features that should be excluded from purging")
	pennantPurgeCmd.Flags().Bool("except-registered", false, "Purge all features except those registered")
	pennantPurgeCmd.Flags().String("store", "", "The store to purge the features from")
	rootCmd.AddCommand(pennantPurgeCmd)
}
