package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "Display the user's identity as known by the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identityCmd).Standalone()

	identityCmd.Flags().BoolS("N", "N", false, "Don't prompt for a password")

	carapace.Gen(identityCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
