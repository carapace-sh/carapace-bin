package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the current status of the record store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_statusCmd).Standalone()

	store_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	storeCmd.AddCommand(store_statusCmd)
}
