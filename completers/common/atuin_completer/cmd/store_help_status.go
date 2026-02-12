package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the current status of the record store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_statusCmd).Standalone()

	store_helpCmd.AddCommand(store_help_statusCmd)
}
