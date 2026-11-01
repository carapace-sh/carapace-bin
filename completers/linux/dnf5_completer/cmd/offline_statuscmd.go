package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "show status of the current offline transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineStatusCmd).Standalone()

	offlineCmd.AddCommand(offlineStatusCmd)
}
