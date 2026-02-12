package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the sync status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_statusCmd).Standalone()

	helpCmd.AddCommand(help_statusCmd)
}
