package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_upgradeLogCmd = &cobra.Command{
	Use:   "log",
	Short: "see logs from offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_upgradeLogCmd).Standalone()

	system_upgradeLogCmd.Flags().Int("number", 0, "boot number")

	system_upgradeCmd.AddCommand(system_upgradeLogCmd)
}
