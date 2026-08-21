package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_unrescueCmd = &cobra.Command{
	Use:   "unrescue",
	Short: "Restore server from rescue mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_unrescueCmd).Standalone()

	serverCmd.AddCommand(server_unrescueCmd)
}
