package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:     "audit",
	Short:   "Checks for known security issues with the installed packages",
	GroupID: "review",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditCmd).Standalone()

	auditCmd.Flags().String("audit-level", "", "Only print advisories with severity greater than or equal to")
	auditCmd.Flags().BoolP("dev", "D", false, "Only audit \"devDependencies\"")
	auditCmd.Flags().Bool("fix", false, "Add overrides to the package.json file in order to force non-vulnerable versions of the dependencies")
	auditCmd.Flags().Bool("ignore-registry-errors", false, "Use exit code 0 if the registry responds with an error")
	auditCmd.Flags().Bool("json", false, "Output audit report in JSON format")
	auditCmd.Flags().Bool("no-optional", false, "Don't audit \"optionalDependencies\"")
	auditCmd.Flags().BoolP("prod", "P", false, "Only audit \"dependencies\" and \"optionalDependencies\"")
	rootCmd.AddCommand(auditCmd)

	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"audit-level": carapace.ActionValues("low", "moderate", "high", "critical"),
	})
}
