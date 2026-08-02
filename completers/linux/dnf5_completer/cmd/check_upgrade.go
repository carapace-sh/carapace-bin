package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var checkUpgradeCmd = &cobra.Command{
	Use:   "check-upgrade [options] [<package-spec>...]",
	Short: "check for available package upgrades",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkUpgradeCmd).Standalone()

	checkUpgradeCmd.Flags().Bool("advisory", false, "Limit to packages specified by advisory")
	checkUpgradeCmd.Flags().String("advisory-severity", "", "Limit to packages specified by advisory severity")
	checkUpgradeCmd.Flags().Bool("bugfix", false, "Limit to packages specified by bugfix advisory")
	checkUpgradeCmd.Flags().String("bz", "", "Limit to packages specified by Bugzilla ID")
	checkUpgradeCmd.Flags().Bool("changelogs", false, "Show changelogs before update")
	checkUpgradeCmd.Flags().String("cve", "", "Limit to packages specified by CVE ID")
	checkUpgradeCmd.Flags().Bool("enhancement", false, "Limit to packages specified by enhancement advisory")
	checkUpgradeCmd.Flags().Bool("json", false, "Request json output format")
	checkUpgradeCmd.Flags().Bool("minimal", false, "Lowest versions fixing advisories")
	checkUpgradeCmd.Flags().Bool("newpackage", false, "Limit to packages specified by newpackage advisory")
	checkUpgradeCmd.Flags().Bool("security", false, "Limit to packages specified by security advisory")

	rootCmd.AddCommand(checkUpgradeCmd)

	carapace.Gen(checkUpgradeCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(checkUpgradeCmd),
	)
}
