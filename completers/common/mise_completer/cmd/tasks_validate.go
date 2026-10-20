// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check task definitions for errors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_validateCmd).Standalone()

	tasks_validateCmd.Flags().Bool("errors-only", false, "Only show errors (skip warnings)")
	tasks_validateCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_validateCmd.Flags().Bool("json", false, "Output validation results in JSON format")
	tasksCmd.AddCommand(tasks_validateCmd)
}
