package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "Suspend server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_suspendCmd).Standalone()

	serverCmd.AddCommand(server_suspendCmd)
}
