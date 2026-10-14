package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var port_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete port(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(port_deleteCmd).Standalone()

	portCmd.AddCommand(port_deleteCmd)
}
