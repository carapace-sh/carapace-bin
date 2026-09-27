package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_unlockCmd).Standalone()

	serverCmd.AddCommand(server_unlockCmd)
}
