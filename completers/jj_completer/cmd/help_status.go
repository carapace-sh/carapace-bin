package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show high-level repo status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_statusCmd).Standalone()

	helpCmd.AddCommand(help_statusCmd)
}
