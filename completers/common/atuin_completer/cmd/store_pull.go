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

	store_pullCmd.Flags().Bool("force", false, "Force pull records")
	store_pullCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	store_pullCmd.Flags().String("page", "100", "Page Size")
	store_pullCmd.Flags().StringP("tag", "t", "", "The tag to push (eg, 'history'). Defaults to all tags")
	storeCmd.AddCommand(store_pullCmd)
}
