package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importKeysCmd = &cobra.Command{
	Use:   "import-keys BACKUP.tar",
	Short: "Import previously backed up Charm account keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importKeysCmd).Standalone()

	importKeysCmd.Flags().BoolP("force-overwrite", "f", false, "overwrite if keys exist; don’t prompt for input")
	rootCmd.AddCommand(importKeysCmd)

	carapace.Gen(importKeysCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
