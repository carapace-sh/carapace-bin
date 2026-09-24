package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_fixedCmd = &cobra.Command{
	Use:   "fixed",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_fixedCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_fixedCmd)
}
