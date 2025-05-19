package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks for modified packages in the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_statusCmd).Standalone()

	storeCmd.AddCommand(store_statusCmd)
}
