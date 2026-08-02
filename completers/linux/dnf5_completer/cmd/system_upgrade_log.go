package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgradeLogCmd = &cobra.Command{
	Use:   "log [options]",
	Short: "show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgradeLogCmd).Standalone()

	systemUpgradeLogCmd.Flags().String("number", "", "Which log to show")

	systemUpgradeCmd.AddCommand(systemUpgradeLogCmd)
}
