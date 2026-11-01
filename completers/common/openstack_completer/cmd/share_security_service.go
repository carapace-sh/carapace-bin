package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_serviceCmd).Standalone()

	share_securityCmd.AddCommand(share_security_serviceCmd)
}
