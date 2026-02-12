package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull records from the remote sync server (one way sync)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pullCmd).Standalone()

	store_pullCmd.Flags().Bool("force", false, "Force push records This will first wipe the local store, and then download all records from the remote")
	store_pullCmd.Flags().BoolP("help", "h", false, "Print help")
	store_pullCmd.Flags().String("page", "", "Page Size How many records to download at once. Defaults to 100")
	store_pullCmd.Flags().StringP("tag", "t", "", "The tag to push (eg, 'history'). Defaults to all tags")
	storeCmd.AddCommand(store_pullCmd)
}
