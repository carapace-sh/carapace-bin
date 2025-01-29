package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_connection_removeCmd = &cobra.Command{
	Use:     "remove [options] NAME",
	Short:   "Delete named destination",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_removeCmd).Standalone()

	system_connection_removeCmd.Flags().BoolP("all", "a", false, "Remove all connections")
	system_connection_removeCmd.Flags().BoolP("force", "f", false, "Ignored: for Docker compatibility")
	system_connection_removeCmd.Flag("force").Hidden = true
	system_connectionCmd.AddCommand(system_connection_removeCmd)
}
