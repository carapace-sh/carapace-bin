package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_stopCmd).Standalone()

	server_stopCmd.Flags().Bool("all-projects", false, "Stop server(s) in another project by name (admin only) (can be specified using the ALL_PROJECTS envvar)")
	serverCmd.AddCommand(server_stopCmd)
}
