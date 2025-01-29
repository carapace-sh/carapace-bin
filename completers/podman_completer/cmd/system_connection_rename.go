package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_connection_renameCmd = &cobra.Command{
	Use:     "rename OLD NEW",
	Short:   "Rename \"old\" to \"new\"",
	Aliases: []string{"mv"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_renameCmd).Standalone()

	system_connectionCmd.AddCommand(system_connection_renameCmd)
}
