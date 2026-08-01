package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var loadKeyCmd = &cobra.Command{
	Use:     "load-key [-nr] [-L keylocation] -a|filesystem",
	Short:   "load encryption key for a dataset",
	GroupID: "encryption",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadKeyCmd).Standalone()

	loadKeyCmd.Flags().StringS("L", "L", "", "use keylocation instead of the keylocation property")
	loadKeyCmd.Flags().BoolS("a", "a", false, "load keys for all encryption roots in all imported pools")
	loadKeyCmd.Flags().BoolS("n", "n", false, "dry-run")
	loadKeyCmd.Flags().BoolS("r", "r", false, "load keys recursively")

	rootCmd.AddCommand(loadKeyCmd)

	carapace.Gen(loadKeyCmd).FlagCompletion(carapace.ActionMap{
		"L": zfs.ActionPropertyValues("keylocation"),
	})

	carapace.Gen(loadKeyCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
