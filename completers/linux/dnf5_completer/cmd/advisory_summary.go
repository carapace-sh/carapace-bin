package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisorySummaryCmd = &cobra.Command{
	Use:   "summary [options] [<advisory-spec>...]",
	Short: "print summary of advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisorySummaryCmd).Standalone()

	advisorySummaryCmd.Flags().String("advisory-severities", "", "Limit to specified severities")
	advisorySummaryCmd.Flags().Bool("all", false, "Show all advisories")
	advisorySummaryCmd.Flags().Bool("available", false, "Show available advisories (default)")
	advisorySummaryCmd.Flags().Bool("bugfix", false, "Show bugfix advisories")
	advisorySummaryCmd.Flags().String("bzs", "", "Limit to specified Bugzilla IDs")
	advisorySummaryCmd.Flags().String("contains-pkgs", "", "Filter by packages")
	advisorySummaryCmd.Flags().String("cves", "", "Limit to specified CVE IDs")
	advisorySummaryCmd.Flags().Bool("enhancement", false, "Show enhancement advisories")
	advisorySummaryCmd.Flags().Bool("installed", false, "Show installed advisories")
	advisorySummaryCmd.Flags().Bool("newpackage", false, "Show newpackage advisories")
	advisorySummaryCmd.Flags().Bool("security", false, "Show security advisories")
	advisorySummaryCmd.Flags().Bool("updates", false, "Show update advisories")
	advisorySummaryCmd.Flags().Bool("with-bz", false, "Include Bugzilla references in output")
	advisorySummaryCmd.Flags().Bool("with-cve", false, "Include CVE references in output")

	advisoryCmd.AddCommand(advisorySummaryCmd)

	carapace.Gen(advisorySummaryCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
	})
}
