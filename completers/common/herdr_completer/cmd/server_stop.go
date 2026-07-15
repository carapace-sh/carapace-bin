package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_stopCmd).Standalone()

	serverCmd.AddCommand(server_stopCmd)
}
