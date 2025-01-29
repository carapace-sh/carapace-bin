package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_connection_listCmd = &cobra.Command{
	Use:     "list [options]",
	Short:   "List destination for the Podman service(s)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_listCmd).Standalone()

	system_connection_listCmd.Flags().StringP("format", "f", "", "Custom Go template for printing connections")
	system_connection_listCmd.Flags().BoolP("quiet", "q", false, "Custom Go template for printing connections")
	system_connectionCmd.AddCommand(system_connection_listCmd)
}
