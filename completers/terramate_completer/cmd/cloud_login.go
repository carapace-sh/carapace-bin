package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login for cloud.terramate.io",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_loginCmd).Standalone()

	cloudCmd.AddCommand(cloud_loginCmd)
}
