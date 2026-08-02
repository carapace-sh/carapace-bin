package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install [options] <package-spec>|@<group-spec>|@<environment-spec>...",
	Aliases: []string{"in"},
	Short:   "install packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("advisories", "", "Include content contained in advisories with specified name")
	installCmd.Flags().String("advisory-severities", "", "Include content contained in advisories with specified severity")
	installCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	installCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	installCmd.Flags().Bool("bugfix", false, "Limit to packages specified by bugfix advisory")
	installCmd.Flags().String("bzs", "", "Include content contained in advisories that fix a Bugzilla ID")
	installCmd.Flags().String("cves", "", "Include content contained in advisories that fix a CVE ID")
	installCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	installCmd.Flags().Bool("enhancement", false, "Limit to packages specified by enhancement advisory")
	installCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	installCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	installCmd.Flags().Bool("newpackage", false, "Limit to packages specified by newpackage advisory")
	installCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	installCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	installCmd.Flags().Bool("security", false, "Limit to packages specified by security advisory")
	installCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	installCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	installCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	installCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
		"destdir":             carapace.ActionDirectories(),
		"store":               carapace.ActionDirectories(),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(installCmd),
	)
}
