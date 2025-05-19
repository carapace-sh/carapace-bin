package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaUpgradeCmd = &cobra.Command{
	Use:   "nova:upgrade",
	Short: "Upgrade Laravel Nova 3 to 4",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaUpgradeCmd).Standalone()

	rootCmd.AddCommand(novaUpgradeCmd)
}
