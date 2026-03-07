package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var advisory_summaryCmd = &cobra.Command{
	Use:   "summary [advisory-spec]...",
	Short: "print a summary count of advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisory_summaryCmd).Standalone()

	advisory_summaryCmd.Flags().String("advisory-severities", "", "include advisories with specified severity")
	advisory_summaryCmd.Flags().Bool("all", false, "show advisories containing any version of installed packages")
	advisory_summaryCmd.Flags().Bool("available", false, "show advisories containing newer versions of installed packages")
	advisory_summaryCmd.Flags().Bool("bugfix", false, "include bugfix advisories")
	advisory_summaryCmd.Flags().String("bzs", "", "include advisories that fix Bugzilla ID")
	advisory_summaryCmd.Flags().String("contains-pkgs", "", "show only advisories containing packages with specified names")
	advisory_summaryCmd.Flags().String("cves", "", "include advisories that fix CVE ID")
	advisory_summaryCmd.Flags().Bool("enhancement", false, "include enhancement advisories")
	advisory_summaryCmd.Flags().Bool("installed", false, "show advisories containing equal and older versions of installed packages")
	advisory_summaryCmd.Flags().Bool("newpackage", false, "include newpackage advisories")
	advisory_summaryCmd.Flags().Bool("security", false, "include security advisories")
	advisory_summaryCmd.Flags().Bool("updates", false, "show advisories containing newer versions of installed packages")
	advisory_summaryCmd.Flags().Bool("with-bz", false, "show only advisories referencing a Bugzilla ticket")
	advisory_summaryCmd.Flags().Bool("with-cve", false, "show only advisories referencing a CVE ticket")
	advisoryCmd.AddCommand(advisory_summaryCmd)

	carapace.Gen(advisory_summaryCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(advisory_summaryCmd),
	)
}
