package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks for modified packages in the store. Returns exit code 0 if the content of the package is the same as it was at the time of unpacking",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_statusCmd).Standalone()

	help_storeCmd.AddCommand(help_store_statusCmd)
}
