package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_backupCmd).Standalone()

	serverCmd.AddCommand(server_backupCmd)
}
