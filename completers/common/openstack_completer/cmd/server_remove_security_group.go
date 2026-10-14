package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_security_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Remove security group from server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_security_groupCmd).Standalone()

	server_remove_securityCmd.AddCommand(server_remove_security_groupCmd)
}
