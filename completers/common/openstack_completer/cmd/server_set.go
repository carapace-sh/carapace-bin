package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set server properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_setCmd).Standalone()

	server_setCmd.Flags().Bool("auto-approve", false, "Allow server state override without asking for confirmation")
	server_setCmd.Flags().String("description", "", "New server description (supported by --os-compute-api-version 2.19 or above)")
	server_setCmd.Flags().String("hostname", "", "Hostname configured for the server in the metadata service.")
	server_setCmd.Flags().String("name", "", "New server name")
	server_setCmd.Flags().Bool("no-password", false, "Clear the admin password for the server from the metadata service; note that this action does not actually change the server password")
	server_setCmd.Flags().String("password", "", "Set the server password.")
	server_setCmd.Flags().String("pinned-availability-zone", "", "Pin the server to the given availability zone.")
	server_setCmd.Flags().String("property", "", "Property to add/change for this server (repeat option to set multiple properties)")
	server_setCmd.Flags().Bool("root-password", false, "==SUPPRESS==")
	server_setCmd.Flags().String("state", "", "New server state.**WARNING** Resetting the state is intended to work around servers stuck in an intermediate state, such as deleting.")
	server_setCmd.Flags().String("tag", "", "Tag for the server.")
	serverCmd.AddCommand(server_setCmd)
}
