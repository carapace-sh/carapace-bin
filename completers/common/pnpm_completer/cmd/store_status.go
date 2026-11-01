package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks for modified packages in the store. Returns exit code 0 if the content of the package is the same as it was at the time of unpacking",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_statusCmd).Standalone()

	store_statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	storeCmd.AddCommand(store_statusCmd)
}
