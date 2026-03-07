package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkUpgradeCmd = &cobra.Command{
	Use:   "check-upgrade [options] [<package-spec>...]",
	Short: "check for available package upgrades",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkUpgradeCmd).Standalone()

	checkUpgradeCmd.Flags().Bool("changelogs", false, "print the package changelogs")
	checkUpgradeCmd.Flags().Bool("json", false, "request JSON output format")
	checkUpgradeCmd.Flags().Bool("minimal", false, "reports the lowest versions of packages that fix advisories")

	rootCmd.AddCommand(checkUpgradeCmd)
}
