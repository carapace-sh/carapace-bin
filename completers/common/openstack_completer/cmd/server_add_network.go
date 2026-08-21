package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Add network to server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_networkCmd).Standalone()

	server_add_networkCmd.Flags().String("tag", "", "Tag for the attached interface.")
	server_addCmd.AddCommand(server_add_networkCmd)
}
