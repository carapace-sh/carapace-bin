package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Manage remote API service destinations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connectionCmd).Standalone()

	systemCmd.AddCommand(system_connectionCmd)
}
