package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_fixedCmd = &cobra.Command{
	Use:   "fixed",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_fixedCmd).Standalone()

	server_addCmd.AddCommand(server_add_fixedCmd)
}
