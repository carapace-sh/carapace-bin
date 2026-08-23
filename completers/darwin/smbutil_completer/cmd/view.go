package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "List resources available on a server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCmd).Standalone()

	viewCmd.Flags().BoolS("A", "A", false, "Authorize only")
	viewCmd.Flags().BoolS("G", "G", false, "Allow guest access")
	viewCmd.Flags().BoolS("N", "N", false, "Don't prompt for a password")
	viewCmd.Flags().BoolS("a", "a", false, "Authorize with anonymous only")
	viewCmd.Flags().BoolS("f", "f", false, "Don't share session")
	viewCmd.Flags().BoolS("g", "g", false, "Authorize with guest only")

	carapace.Gen(viewCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
