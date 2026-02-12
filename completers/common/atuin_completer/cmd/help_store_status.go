package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the current status of the record store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_statusCmd).Standalone()

	help_storeCmd.AddCommand(help_store_statusCmd)
}
