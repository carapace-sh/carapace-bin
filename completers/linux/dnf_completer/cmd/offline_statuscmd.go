package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show status of current offline transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline_statusCmd).Standalone()

	offlineCmd.AddCommand(offline_statusCmd)
}
