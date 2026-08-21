package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_startCmd).Standalone()

	server_startCmd.Flags().Bool("all-projects", false, "Start server(s) in another project by name (admin only) (can be specified using the ALL_PROJECTS envvar)")
	serverCmd.AddCommand(server_startCmd)
}
