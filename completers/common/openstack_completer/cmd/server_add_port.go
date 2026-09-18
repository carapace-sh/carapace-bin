package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_portCmd = &cobra.Command{
	Use:   "port",
	Short: "Add port to server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_portCmd).Standalone()

	server_add_portCmd.Flags().String("tag", "", "Tag for the attached interface (supported by --os-compute-api-version 2.49 or later)")
	server_addCmd.AddCommand(server_add_portCmd)
}
