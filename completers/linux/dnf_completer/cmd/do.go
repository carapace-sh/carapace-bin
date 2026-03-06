package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do [options] [arguments]",
	Short: "execute a multi-action transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doCmd).Standalone()

	doCmd.Flags().String("action", "", "action to be done (install, remove, upgrade, downgrade, reinstall)")
	doCmd.Flags().String("type", "", "type of the following items (auto, package, group)")
	doCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	doCmd.Flags().Bool("skip-broken", false, "resolve dependency problems by removing problematic packages")
	doCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not available")
	doCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	doCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	doCmd.Flags().String("installed-from-repo", "", "limit the operation to packages installed from the given repository")
	doCmd.Flags().String("from-repo", "", "limit the operation to the given repository")
	doCmd.Flags().String("from-vendor", "", "limit the operation to packages from the given vendor")
	doCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	doCmd.Flags().Bool("offline", false, "store the transaction to be performed offline")
	doCmd.Flags().StringSlice("advisories", []string{}, "limit the operation to packages matching the given advisory")
	doCmd.Flags().StringSlice("advisory-severities", []string{}, "limit the operation to packages matching the given advisory severity")
	doCmd.Flags().StringSlice("bzs", []string{}, "limit the operation to packages matching the given bugzilla")
	doCmd.Flags().StringSlice("cves", []string{}, "limit the operation to packages matching the given cve")
	doCmd.Flags().Bool("security", false, "limit the operation to packages matching the given security criteria")
	doCmd.Flags().Bool("bugfix", false, "limit to packages that fix a bugfix advisory")
	doCmd.Flags().Bool("enhancement", false, "limit to packages that fix an enhancement advisory")
	doCmd.Flags().Bool("newpackage", false, "limit to packages that fix a newpackage advisory")

	rootCmd.AddCommand(doCmd)

	carapace.Gen(doCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(doCmd),
	)
}
