package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stash entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_listCmd).Standalone()

	stash_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stashCmd.AddCommand(stash_listCmd)
}
