package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push all records to the remote sync server (one way sync)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pushCmd).Standalone()

	store_pushCmd.Flags().Bool("force", false, "Force push records")
	store_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	store_pushCmd.Flags().String("host", "", "The host to push, in the form of a UUID host ID. Defaults to the current host")
	store_pushCmd.Flags().String("page", "100", "Page Size")
	store_pushCmd.Flags().StringP("tag", "t", "", "The tag to push (eg, 'history'). Defaults to all tags")
	storeCmd.AddCommand(store_pushCmd)
}
