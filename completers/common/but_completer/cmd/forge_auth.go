package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forge_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticat with your forge provider (at the moment, only GitHub is supported)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forge_authCmd).Standalone()

	forge_authCmd.Flags().BoolP("help", "h", false, "Print help")
	forgeCmd.AddCommand(forge_authCmd)
}
