package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_devStoreSetCmd = &cobra.Command{
	Use:   "dev-store-set",
	Short: "Set a key/value pair during development",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_devStoreSetCmd).Standalone()

	debug_devStoreSetCmd.Flags().Bool("danger", false, "acknowledge the danger of this operation")
	debugCmd.AddCommand(debug_devStoreSetCmd)
}
