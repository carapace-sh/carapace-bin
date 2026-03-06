package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var advisoryCmd = &cobra.Command{
	Use:   "advisory <subcommand> [options] [<advisory-spec>...]",
	Short: "manage advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var advisoryListCmd = &cobra.Command{
	Use:   "list [advisory-spec]...",
	Short: "list advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var advisoryInfoCmd = &cobra.Command{
	Use:   "info [advisory-spec]...",
	Short: "print detailed information about advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var advisorySummaryCmd = &cobra.Command{
	Use:   "summary [advisory-spec]...",
	Short: "print a summary count of advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisoryCmd).Standalone()

	advisoryCmd.AddCommand(advisoryListCmd)
	advisoryCmd.AddCommand(advisoryInfoCmd)
	advisoryCmd.AddCommand(advisorySummaryCmd)

	for _, subcmd := range []*cobra.Command{advisoryListCmd, advisoryInfoCmd, advisorySummaryCmd} {
		subcmd.Flags().Bool("all", false, "show advisories containing any version of installed packages")
		subcmd.Flags().Bool("available", false, "show advisories containing newer versions of installed packages")
		subcmd.Flags().Bool("installed", false, "show advisories containing equal and older versions of installed packages")
		subcmd.Flags().Bool("updates", false, "show advisories containing newer versions of installed packages")
		subcmd.Flags().String("contains-pkgs", "", "show only advisories containing packages with specified names")
		subcmd.Flags().Bool("security", false, "include security advisories")
		subcmd.Flags().Bool("bugfix", false, "include bugfix advisories")
		subcmd.Flags().Bool("enhancement", false, "include enhancement advisories")
		subcmd.Flags().Bool("newpackage", false, "include newpackage advisories")
		subcmd.Flags().String("advisory-severities", "", "include advisories with specified severity")
		subcmd.Flags().String("bzs", "", "include advisories that fix Bugzilla ID")
		subcmd.Flags().String("cves", "", "include advisories that fix CVE ID")
		subcmd.Flags().Bool("with-bz", false, "show only advisories referencing a Bugzilla ticket")
		subcmd.Flags().Bool("with-cve", false, "show only advisories referencing a CVE ticket")
	}

	// Additional flags for list and info
	advisoryListCmd.Flags().Bool("json", false, "request JSON output format")
	advisoryInfoCmd.Flags().Bool("json", false, "request JSON output format")

	rootCmd.AddCommand(advisoryCmd)

	carapace.Gen(advisoryListCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(advisoryListCmd),
	)
	carapace.Gen(advisoryInfoCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(advisoryInfoCmd),
	)
	carapace.Gen(advisorySummaryCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(advisorySummaryCmd),
	)
}
