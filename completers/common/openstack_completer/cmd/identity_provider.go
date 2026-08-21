package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identity_providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_providerCmd).Standalone()

	identityCmd.AddCommand(identity_providerCmd)
}
