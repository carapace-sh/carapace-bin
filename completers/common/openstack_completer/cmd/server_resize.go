package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Scale server to a new flavor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_resizeCmd).Standalone()

	server_resizeCmd.Flags().Bool("confirm", false, "**Deprecated** Confirm server resize is complete.")
	server_resizeCmd.Flags().String("flavor", "", "Resize server to specified flavor")
	server_resizeCmd.Flags().Bool("revert", false, "**Deprecated** Restore server state before resize.")
	server_resizeCmd.Flags().Bool("wait", false, "Wait for resize to complete")
	serverCmd.AddCommand(server_resizeCmd)
}
