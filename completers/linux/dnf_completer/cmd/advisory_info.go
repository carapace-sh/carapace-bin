package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var advisory_infoCmd = &cobra.Command{
	Use:   "info [advisory-spec]...",
	Short: "print detailed information about advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisory_infoCmd).Standalone()

	advisory_infoCmd.Flags().String("advisory-severities", "", "include advisories with specified severity")
	advisory_infoCmd.Flags().Bool("all", false, "show advisories containing any version of installed packages")
	advisory_infoCmd.Flags().Bool("available", false, "show advisories containing newer versions of installed packages")
	advisory_infoCmd.Flags().Bool("bugfix", false, "include bugfix advisories")
	advisory_infoCmd.Flags().String("bzs", "", "include advisories that fix Bugzilla ID")
	advisory_infoCmd.Flags().String("contains-pkgs", "", "show only advisories containing packages with specified names")
	advisory_infoCmd.Flags().String("cves", "", "include advisories that fix CVE ID")
	advisory_infoCmd.Flags().Bool("enhancement", false, "include enhancement advisories")
	advisory_infoCmd.Flags().Bool("installed", false, "show advisories containing equal and older versions of installed packages")
	advisory_infoCmd.Flags().Bool("json", false, "request JSON output format")
	advisory_infoCmd.Flags().Bool("newpackage", false, "include newpackage advisories")
	advisory_infoCmd.Flags().Bool("security", false, "include security advisories")
	advisory_infoCmd.Flags().Bool("updates", false, "show advisories containing newer versions of installed packages")
	advisory_infoCmd.Flags().Bool("with-bz", false, "show only advisories referencing a Bugzilla ticket")
	advisory_infoCmd.Flags().Bool("with-cve", false, "show only advisories referencing a CVE ticket")
	advisoryCmd.AddCommand(advisory_infoCmd)

	carapace.Gen(advisory_infoCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(advisory_infoCmd),
	)
}
