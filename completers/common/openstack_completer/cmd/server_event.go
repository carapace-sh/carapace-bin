package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_eventCmd = &cobra.Command{
	Use:   "event",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_eventCmd).Standalone()

	serverCmd.AddCommand(server_eventCmd)
}
