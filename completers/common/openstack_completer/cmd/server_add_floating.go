package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_floatingCmd = &cobra.Command{
	Use:   "floating",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_floatingCmd).Standalone()

	server_addCmd.AddCommand(server_add_floatingCmd)
}
