package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisoryListCmd = &cobra.Command{
	Use:   "list [options] [<advisory-spec>...]",
	Short: "list advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisoryListCmd).Standalone()

	advisoryListCmd.Flags().String("advisory-severities", "", "Limit to specified severities")
	advisoryListCmd.Flags().Bool("all", false, "Show all advisories")
	advisoryListCmd.Flags().Bool("available", false, "Show available advisories (default)")
	advisoryListCmd.Flags().Bool("bugfix", false, "Show bugfix advisories")
	advisoryListCmd.Flags().String("bzs", "", "Limit to specified Bugzilla IDs")
	advisoryListCmd.Flags().String("contains-pkgs", "", "Filter by packages")
	advisoryListCmd.Flags().String("cves", "", "Limit to specified CVE IDs")
	advisoryListCmd.Flags().Bool("enhancement", false, "Show enhancement advisories")
	advisoryListCmd.Flags().Bool("installed", false, "Show installed advisories")
	advisoryListCmd.Flags().Bool("json", false, "Request json output format")
	advisoryListCmd.Flags().Bool("newpackage", false, "Show newpackage advisories")
	advisoryListCmd.Flags().Bool("security", false, "Show security advisories")
	advisoryListCmd.Flags().Bool("updates", false, "Show update advisories")
	advisoryListCmd.Flags().Bool("with-bz", false, "Include Bugzilla references in output")
	advisoryListCmd.Flags().Bool("with-cve", false, "Include CVE references in output")

	advisoryCmd.AddCommand(advisoryListCmd)

	carapace.Gen(advisoryListCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
	})
}
