package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgradeLogCmd = &cobra.Command{
	Use:   "log [options]",
	Short: "show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgradeLogCmd).Standalone()

	offlineUpgradeLogCmd.Flags().String("number", "", "Which log to show")

	offlineUpgradeCmd.AddCommand(offlineUpgradeLogCmd)
}
