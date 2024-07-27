package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireUpgradeCmd = &cobra.Command{
	Use:   "livewire:upgrade [--run-only [RUN-ONLY]]",
	Short: "Interactive upgrade helper to migrate from v2 to v3",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireUpgradeCmd).Standalone()

	livewireUpgradeCmd.Flags().String("run-only", "", "")
	rootCmd.AddCommand(livewireUpgradeCmd)
}
