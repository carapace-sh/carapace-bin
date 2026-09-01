package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Checks for known security issues with the installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditCmd).Standalone()

	auditCmd.Flags().String("audit-level", "", "Only print advisories with severity greater than or equal to this level")
	auditCmd.Flags().BoolP("dev", "D", false, "Only audit \"devDependencies\"")
	auditCmd.Flags().String("fix", "", "Fix the audited vulnerabilities using the specified method: \"override\" or \"update\". \"override\" adds overrides to `pnpm-workspace.yaml` to force non-vulnerable versions; \"update\" re-resolves the lockfile to non-vulnerable versions. Defaults to \"override\" when no method is given")
	auditCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	auditCmd.Flags().StringSlice("ignore", nil, "Ignore a vulnerability by its GitHub advisory ID (e.g. GHSA-xxxx-xxxx-xxxx). May be repeated")
	auditCmd.Flags().Bool("ignore-registry-errors", false, "Use exit code 0 if the registry responds with an error")
	auditCmd.Flags().Bool("ignore-unfixable", false, "Ignore all vulnerabilities for which no fix exists")
	auditCmd.Flags().BoolP("interactive", "i", false, "Show vulnerabilities and select which ones to fix interactively")
	auditCmd.Flags().Bool("json", false, "Output audit report in JSON format")
	auditCmd.Flags().Bool("no-optional", false, "Don't audit \"optionalDependencies\"")
	auditCmd.Flags().Bool("optional", false, "Include \"optionalDependencies\"")
	auditCmd.Flags().BoolP("prod", "P", false, "Only audit \"dependencies\" and \"optionalDependencies\"")
	auditCmd.Flags().Bool("production", false, "Only audit \"dependencies\" and \"optionalDependencies\"")

	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"audit-level": carapace.ActionValues("info", "low", "moderate", "high", "critical"),
	})

	rootCmd.AddCommand(auditCmd)
}
