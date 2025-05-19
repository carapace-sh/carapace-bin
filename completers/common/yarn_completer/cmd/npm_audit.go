package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var npm_auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Perform a vulnerability audit against the installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_auditCmd).Standalone()

	npm_auditCmd.Flags().BoolP("all", "A", false, "Audit dependencies from all workspaces")
	npm_auditCmd.Flags().String("environment", "", "Which environments to cover")
	npm_auditCmd.Flags().String("exclude", "", "Array of glob patterns of packages to exclude from audit")
	npm_auditCmd.Flags().String("ignore", "", "Array of glob patterns of advisory ID's to ignore in the audit report")
	npm_auditCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	npm_auditCmd.Flags().BoolP("recursive", "R", false, "Audit transitive dependencies as well")
	npm_auditCmd.Flags().String("severity", "", "Minimal severity requested for packages to be displayed")
	npmCmd.AddCommand(npm_auditCmd)

	// TODO environment
	carapace.Gen(npm_auditCmd).FlagCompletion(carapace.ActionMap{
		"severity": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionStyledValues(
				"info", style.Carapace.LogLevelInfo,
				"low", style.Carapace.LogLevelDebug,
				"moderate", style.Carapace.LogLevelError,
				"high", style.Carapace.LogLevelFatal,
				"critical", style.Carapace.LogLevelCritical,
			)
		}),
	})
}
