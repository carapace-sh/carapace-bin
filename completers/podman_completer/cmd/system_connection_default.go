package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_connection_defaultCmd = &cobra.Command{
	Use:   "default NAME",
	Short: "Set named destination as default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_defaultCmd).Standalone()

	system_connectionCmd.AddCommand(system_connection_defaultCmd)
}
