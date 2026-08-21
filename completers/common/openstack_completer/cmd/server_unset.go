package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset server properties and tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_unsetCmd).Standalone()

	server_unsetCmd.Flags().Bool("all-properties", false, "Remove all properties")
	server_unsetCmd.Flags().Bool("all-tags", false, "Remove all tags (supported by --os-compute-api-version 2.26 or above)")
	server_unsetCmd.Flags().Bool("description", false, "Unset server description (supported by --os-compute-api-version 2.19 or above)")
	server_unsetCmd.Flags().Bool("pinned-availability-zone", false, "Unpin the server from its availability zone (supported by --os-compute-api-version 2.104 or above)")
	server_unsetCmd.Flags().String("property", "", "Property key to remove from server (repeat option to remove multiple values)")
	server_unsetCmd.Flags().String("tag", "", "Tag to remove from the server.")
	serverCmd.AddCommand(server_unsetCmd)
}
