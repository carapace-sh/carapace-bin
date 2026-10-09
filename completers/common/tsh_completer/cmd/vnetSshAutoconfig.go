package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vnetSshAutoconfigCmd = &cobra.Command{
	Use:   "vnet-ssh-autoconfig",
	Short: "Automatically include VNet's generated OpenSSH-compatible config file in ~/.ssh/config.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vnetSshAutoconfigCmd).Standalone()

	rootCmd.AddCommand(vnetSshAutoconfigCmd)
}
