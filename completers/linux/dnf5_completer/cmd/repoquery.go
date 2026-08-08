package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repoqueryCmd = &cobra.Command{
	Use:   "repoquery [options] [<package-spec>...]",
	Short: "search for packages in repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoqueryCmd).Standalone()

	repoqueryCmd.Flags().String("advisories", "", "Include content contained in advisories with specified name")
	repoqueryCmd.Flags().String("advisory-severities", "", "Include content contained in advisories with specified severity")
	repoqueryCmd.Flags().Bool("all", false, "Query all packages")
	repoqueryCmd.Flags().String("arch", "", "Limit by architectures")
	repoqueryCmd.Flags().Bool("available", false, "Query available packages (default)")
	repoqueryCmd.Flags().Bool("bugfix", false, "Limit to packages specified by bugfix advisory")
	repoqueryCmd.Flags().String("bzs", "", "Include content contained in advisories that fix a Bugzilla ID")
	repoqueryCmd.Flags().Bool("changelogs", false, "Display changelogs")
	repoqueryCmd.Flags().Bool("conflicts", false, "Like --queryformat=%{conflicts}")
	repoqueryCmd.Flags().String("cves", "", "Include content contained in advisories that fix a CVE ID")
	repoqueryCmd.Flags().Bool("depends", false, "Like --queryformat=%{depends}")
	repoqueryCmd.Flags().Bool("disable-modular-filtering", false, "Include inactive module streams")
	repoqueryCmd.Flags().Bool("duplicates", false, "Duplicate installed packages")
	repoqueryCmd.Flags().Bool("enhancement", false, "Limit to packages specified by enhancement advisory")
	repoqueryCmd.Flags().Bool("enhances", false, "Like --queryformat=%{enhances}")
	repoqueryCmd.Flags().Bool("exactdeps", false, "Exact requires matching")
	repoqueryCmd.Flags().Bool("extras", false, "Installed but not in any available repo")
	repoqueryCmd.Flags().String("file", "", "Limit by owned files")
	repoqueryCmd.Flags().String("from-repo", "", "Limit to packages from specific repos")
	repoqueryCmd.Flags().BoolP("info", "i", false, "Show detailed package information")
	repoqueryCmd.Flags().Bool("installed", false, "Query installed packages")
	repoqueryCmd.Flags().Bool("installonly", false, "Installed installonly packages")
	repoqueryCmd.Flags().Int32("latest-limit", 0, "Limit to N latest packages per name.arch")
	repoqueryCmd.Flags().Bool("leaves", false, "Installed packages not required by other installed packages")
	repoqueryCmd.Flags().Bool("location", false, "Like --queryformat=%{location}")
	repoqueryCmd.Flags().Bool("newpackage", false, "Limit to packages specified by newpackage advisory")
	repoqueryCmd.Flags().Bool("obsoletes", false, "Like --queryformat=%{obsoletes}")
	repoqueryCmd.Flags().String("providers-of", "", "Providers of selected attribute")
	repoqueryCmd.Flags().Bool("provides", false, "Like --queryformat=%{provides}")
	repoqueryCmd.Flags().String("queryformat", "", "Custom output format")
	repoqueryCmd.Flags().Bool("querytags", false, "Display available tags for --queryformat")
	repoqueryCmd.Flags().Bool("recent", false, "Recently changed packages")
	repoqueryCmd.Flags().Bool("recommends", false, "Like --queryformat=%{recommends}")
	repoqueryCmd.Flags().Bool("recursive", false, "Recursive query")
	repoqueryCmd.Flags().Bool("requires", false, "Like --queryformat=%{requires}")
	repoqueryCmd.Flags().Bool("requires-pre", false, "Like --queryformat=%{requires_pre}")
	repoqueryCmd.Flags().Bool("security", false, "Limit to packages specified by security advisory")
	repoqueryCmd.Flags().Bool("sourcerpm", false, "Like --queryformat=%{sourcerpm}")
	repoqueryCmd.Flags().Bool("srpm", false, "Output source RPMs instead")
	repoqueryCmd.Flags().Bool("suggests", false, "Like --queryformat=%{suggests}")
	repoqueryCmd.Flags().Bool("supplements", false, "Like --queryformat=%{supplements}")
	repoqueryCmd.Flags().Bool("unneeded", false, "Unneeded dependencies")
	repoqueryCmd.Flags().Bool("upgrades", false, "Available packages that upgrade installed ones")
	repoqueryCmd.Flags().Bool("userinstalled", false, "Packages not installed as dependencies")
	repoqueryCmd.Flags().String("whatconflicts", "", "Packages that conflict with any")
	repoqueryCmd.Flags().String("whatdepends", "", "Packages that require/enhance/recommend/suggest/supplement")
	repoqueryCmd.Flags().String("whatenhances", "", "Packages that enhance any")
	repoqueryCmd.Flags().String("whatobsoletes", "", "Packages that obsolete any")
	repoqueryCmd.Flags().String("whatprovides", "", "Packages that provide any")
	repoqueryCmd.Flags().String("whatrecommends", "", "Packages that recommend any")
	repoqueryCmd.Flags().String("whatrequires", "", "Packages that require any")
	repoqueryCmd.Flags().String("whatsuggests", "", "Packages that suggest any")
	repoqueryCmd.Flags().String("whatsupplements", "", "Packages that supplement any")

	rootCmd.AddCommand(repoqueryCmd)

	carapace.Gen(repoqueryCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
		"providers-of":        carapace.ActionValues("conflicts", "depends", "enhances", "obsoletes", "provides", "recommends", "requires", "requires_pre", "suggests", "supplements"),
	})

	carapace.Gen(repoqueryCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(repoqueryCmd),
	)
}
