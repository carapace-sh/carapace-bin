package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current status of database logging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_statusCmd).Standalone()

	logsCmd.AddCommand(logs_statusCmd)
}
