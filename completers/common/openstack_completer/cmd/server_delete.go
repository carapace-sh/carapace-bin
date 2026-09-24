package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_deleteCmd).Standalone()

	server_deleteCmd.Flags().Bool("all-projects", false, "Delete server(s) in another project by name (admin only)(can be specified using the ALL_PROJECTS envvar)")
	server_deleteCmd.Flags().Bool("force", false, "Force delete server(s)")
	server_deleteCmd.Flags().Bool("wait", false, "Wait for delete to complete")
	serverCmd.AddCommand(server_deleteCmd)
}
