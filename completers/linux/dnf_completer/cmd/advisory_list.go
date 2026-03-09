package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var advisory_listCmd = &cobra.Command{
	Use:   "list [advisory-spec]...",
	Short: "list advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisory_listCmd).Standalone()

	advisory_listCmd.Flags().String("advisory-severities", "", "include advisories with specified severity")
	advisory_listCmd.Flags().Bool("all", false, "show advisories containing any version of installed packages")
	advisory_listCmd.Flags().Bool("available", false, "show advisories containing newer versions of installed packages")
	advisory_listCmd.Flags().Bool("bugfix", false, "include bugfix advisories")
	advisory_listCmd.Flags().String("bzs", "", "include advisories that fix Bugzilla ID")
	advisory_listCmd.Flags().String("contains-pkgs", "", "show only advisories containing packages with specified names")
	advisory_listCmd.Flags().String("cves", "", "include advisories that fix CVE ID")
	advisory_listCmd.Flags().Bool("enhancement", false, "include enhancement advisories")
	advisory_listCmd.Flags().Bool("installed", false, "show advisories containing equal and older versions of installed packages")
	advisory_listCmd.Flags().Bool("json", false, "request JSON output format")
	advisory_listCmd.Flags().Bool("newpackage", false, "include newpackage advisories")
	advisory_listCmd.Flags().Bool("security", false, "include security advisories")
	advisory_listCmd.Flags().Bool("updates", false, "show advisories containing newer versions of installed packages")
	advisory_listCmd.Flags().Bool("with-bz", false, "show only advisories referencing a Bugzilla ticket")
	advisory_listCmd.Flags().Bool("with-cve", false, "show only advisories referencing a CVE ticket")
	advisoryCmd.AddCommand(advisory_listCmd)

	carapace.Gen(advisory_listCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(advisory_listCmd),
	)
}
