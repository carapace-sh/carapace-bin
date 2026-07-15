package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_reloadConfigCmd = &cobra.Command{
	Use:   "reload-config",
	Short: "Reload config in the running server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_reloadConfigCmd).Standalone()

	serverCmd.AddCommand(server_reloadConfigCmd)
}
