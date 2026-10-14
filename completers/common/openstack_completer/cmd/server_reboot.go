package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Perform a hard or soft server reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_rebootCmd).Standalone()

	server_rebootCmd.Flags().Bool("hard", false, "Perform a hard reboot")
	server_rebootCmd.Flags().Bool("soft", false, "Perform a soft reboot")
	server_rebootCmd.Flags().Bool("wait", false, "Wait for reboot to complete")
	serverCmd.AddCommand(server_rebootCmd)
}
