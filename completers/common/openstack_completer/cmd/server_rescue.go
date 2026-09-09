package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_rescueCmd = &cobra.Command{
	Use:   "rescue",
	Short: "Put server in rescue mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_rescueCmd).Standalone()

	server_rescueCmd.Flags().String("image", "", "Image (name or ID) to use for the rescue mode (defaults to the currently used one)")
	server_rescueCmd.Flags().String("password", "", "Set the password on the rescued instance (requires cloud support)")
	serverCmd.AddCommand(server_rescueCmd)
}
