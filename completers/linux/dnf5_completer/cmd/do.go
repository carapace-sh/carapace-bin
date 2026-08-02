package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do [options] <items>...",
	Short: "execute multiple actions in a single transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doCmd).Standalone()

	doCmd.Flags().String("action", "", "Action to be done on the following items")
	doCmd.Flags().String("advisories", "", "Include content contained in advisories with specified name")
	doCmd.Flags().String("advisory-severities", "", "Include content contained in advisories with specified severity")
	doCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	doCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	doCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	doCmd.Flags().Bool("bugfix", false, "Limit to packages specified by bugfix advisory")
	doCmd.Flags().String("bzs", "", "Include content contained in advisories that fix a Bugzilla ID")
	doCmd.Flags().String("cves", "", "Include content contained in advisories that fix a CVE ID")
	doCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	doCmd.Flags().Bool("enhancement", false, "Limit to packages specified by enhancement advisory")
	doCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	doCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	doCmd.Flags().Bool("newpackage", false, "Limit to packages specified by newpackage advisory")
	doCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	doCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	doCmd.Flags().Bool("security", false, "Limit to packages specified by security advisory")
	doCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	doCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	doCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	doCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")
	doCmd.Flags().String("type", "", "Type of the following items")

	rootCmd.AddCommand(doCmd)

	carapace.Gen(doCmd).FlagCompletion(carapace.ActionMap{
		"action":              carapace.ActionValues("install", "remove", "upgrade", "downgrade", "reinstall"),
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
		"store":               carapace.ActionDirectories(),
		"type":                carapace.ActionValues("auto", "package", "group", "environment"),
	})

	carapace.Gen(doCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(doCmd),
	)
}
