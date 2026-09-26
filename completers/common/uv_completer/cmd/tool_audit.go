package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tool_auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit installed tools and their dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_auditCmd).Standalone()

	tool_auditCmd.Flags().Bool("all", false, "Audit all installed tools")
	tool_auditCmd.Flags().StringSlice("ignore", nil, "Ignore a vulnerability by ID")
	tool_auditCmd.Flags().StringSlice("ignore-until-fixed", nil, "Ignore a vulnerability by ID, but only while no fix is available")
	tool_auditCmd.Flags().String("output-format", "text", "Select the output format")
	tool_auditCmd.Flags().String("service-format", "osv", "The service format to use for vulnerability lookups")
	tool_auditCmd.Flags().String("service-url", "", "The URL to vulnerability service API endpoint")
	toolCmd.AddCommand(tool_auditCmd)
	carapace.Gen(tool_auditCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the result in a human-readable format",
			"json", "Display the result in JSON format",
			"sarif", "Display the result in SARIF format",
		),
		"service-format": carapace.ActionValues("osv"),
	})
}
