package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_securityCmd = &cobra.Command{
	Use:   "security",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_securityCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_securityCmd)
}
