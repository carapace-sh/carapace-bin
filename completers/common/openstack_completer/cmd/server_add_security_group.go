package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_security_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Add security group(s) to server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_security_groupCmd).Standalone()

	server_add_securityCmd.AddCommand(server_add_security_groupCmd)
}
