package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migrationCmd).Standalone()

	serverCmd.AddCommand(server_migrationCmd)
}
