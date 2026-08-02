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

	checkUpgradeCmd.Flags().String("advisories", "", "Include content contained in advisories with specified name")
	checkUpgradeCmd.Flags().String("advisory-severities", "", "Include content contained in advisories with specified severity")
	checkUpgradeCmd.Flags().Bool("bugfix", false, "Limit to packages specified by bugfix advisory")
	checkUpgradeCmd.Flags().String("bzs", "", "Include content contained in advisories that fix a Bugzilla ID")
	checkUpgradeCmd.Flags().Bool("changelogs", false, "Show changelogs before update")
	checkUpgradeCmd.Flags().String("cves", "", "Include content contained in advisories that fix a CVE ID")
	checkUpgradeCmd.Flags().Bool("enhancement", false, "Limit to packages specified by enhancement advisory")
	checkUpgradeCmd.Flags().Bool("json", false, "Request json output format")
	checkUpgradeCmd.Flags().Bool("minimal", false, "Lowest versions fixing advisories")
	checkUpgradeCmd.Flags().Bool("newpackage", false, "Limit to packages specified by newpackage advisory")
	checkUpgradeCmd.Flags().Bool("security", false, "Limit to packages specified by security advisory")

	rootCmd.AddCommand(checkUpgradeCmd)

	carapace.Gen(checkUpgradeCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
	})

	carapace.Gen(checkUpgradeCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(checkUpgradeCmd),
	)
}
