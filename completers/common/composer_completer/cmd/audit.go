package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Checks for security vulnerability advisories for installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditCmd).Standalone()

	auditCmd.Flags().String("abandoned", "", "Behavior on abandoned packages. Must be \"ignore\", \"report\", or \"fail\".")
	auditCmd.Flags().StringP("format", "f", "", "Output format. Must be \"table\", \"plain\", \"json\", or \"summary\".")
	auditCmd.Flags().StringArray("ignore-severity", nil, "Ignore advisories of a certain severity level.")
	auditCmd.Flags().Bool("ignore-unreachable", false, "Ignore repositories that are unreachable or return a non-200 status code.")
	auditCmd.Flags().Bool("locked", false, "Audit based on the lock file instead of the installed packages.")
	auditCmd.Flags().Bool("no-dev", false, "Disables auditing of require-dev packages.")

	rootCmd.AddCommand(auditCmd)

	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"abandoned":       carapace.ActionValues("fail", "ignore", "report"),
		"format":          carapace.ActionValues("json", "plain", "summary", "table"),
		"ignore-severity": carapace.ActionValues("critical", "high", "low", "medium"),
	})
}
