package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var step_evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "[experimental] Evaluate a template expression",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_evalCmd).Standalone()

	step_evalCmd.Flags().Bool("dry-run", false, "Show template variables and expanded result")
	step_evalCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_evalCmd)
}
