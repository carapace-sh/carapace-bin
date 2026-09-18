package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_service_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more security services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_service_deleteCmd).Standalone()

	share_security_serviceCmd.AddCommand(share_security_service_deleteCmd)
}
