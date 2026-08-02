package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade [options] [<package-spec>|@<group-spec>|@<environment-spec>...]",
	Short: "upgrade software",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("advisories", "", "Include content contained in advisories with specified name")
	upgradeCmd.Flags().String("advisory-severities", "", "Include content contained in advisories with specified severity")
	upgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	upgradeCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	upgradeCmd.Flags().Bool("bugfix", false, "Limit to packages specified by bugfix advisory")
	upgradeCmd.Flags().String("bzs", "", "Include content contained in advisories that fix a Bugzilla ID")
	upgradeCmd.Flags().String("cves", "", "Include content contained in advisories that fix a CVE ID")
	upgradeCmd.Flags().String("destdir", "", "Set directory used for downloading packages to")
	upgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	upgradeCmd.Flags().Bool("enhancement", false, "Limit to packages specified by enhancement advisory")
	upgradeCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	upgradeCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	upgradeCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	upgradeCmd.Flags().Bool("minimal", false, "Upgrade only to lowest versions that fix advisories")
	upgradeCmd.Flags().Bool("newpackage", false, "Limit to packages specified by newpackage advisory")
	upgradeCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	upgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	upgradeCmd.Flags().Bool("security", false, "Limit to packages specified by security advisory")
	upgradeCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	upgradeCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	upgradeCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
		"destdir":             carapace.ActionDirectories(),
		"store":               carapace.ActionDirectories(),
	})

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(upgradeCmd),
	)
}
