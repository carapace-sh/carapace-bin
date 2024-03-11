package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "cloud information status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_infoCmd).Standalone()

	cloudCmd.AddCommand(cloud_infoCmd)
}
