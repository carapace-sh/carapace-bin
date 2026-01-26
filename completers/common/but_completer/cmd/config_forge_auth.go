package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_forge_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your forge provider (currently only GitHub is supported)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_forge_authCmd).Standalone()

	config_forge_authCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_forgeCmd.AddCommand(config_forge_authCmd)
}
