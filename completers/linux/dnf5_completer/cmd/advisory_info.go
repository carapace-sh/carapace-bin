package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisoryInfoCmd = &cobra.Command{
	Use:   "info [options] [<advisory-spec>...]",
	Short: "print details about advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisoryInfoCmd).Standalone()

	advisoryInfoCmd.Flags().String("advisory-severities", "", "Limit to specified severities")
	advisoryInfoCmd.Flags().Bool("all", false, "Show all advisories")
	advisoryInfoCmd.Flags().Bool("available", false, "Show available advisories (default)")
	advisoryInfoCmd.Flags().Bool("bugfix", false, "Show bugfix advisories")
	advisoryInfoCmd.Flags().String("bzs", "", "Limit to specified Bugzilla IDs")
	advisoryInfoCmd.Flags().String("contains-pkgs", "", "Filter by packages")
	advisoryInfoCmd.Flags().String("cves", "", "Limit to specified CVE IDs")
	advisoryInfoCmd.Flags().Bool("enhancement", false, "Show enhancement advisories")
	advisoryInfoCmd.Flags().Bool("installed", false, "Show installed advisories")
	advisoryInfoCmd.Flags().Bool("json", false, "Request json output format")
	advisoryInfoCmd.Flags().Bool("newpackage", false, "Show newpackage advisories")
	advisoryInfoCmd.Flags().Bool("security", false, "Show security advisories")
	advisoryInfoCmd.Flags().Bool("updates", false, "Show update advisories")
	advisoryInfoCmd.Flags().Bool("with-bz", false, "Include Bugzilla references in output")
	advisoryInfoCmd.Flags().Bool("with-cve", false, "Include CVE references in output")

	advisoryCmd.AddCommand(advisoryInfoCmd)

	carapace.Gen(advisoryInfoCmd).FlagCompletion(carapace.ActionMap{
		"advisory-severities": carapace.ActionValues("critical", "important", "moderate", "low", "none"),
	})
}
