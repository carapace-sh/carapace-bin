package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var determineTestRunnersCmd = &cobra.Command{
	Use:     "determine-test-runners",
	Short:   "Determines the runners used to test formulae or their dependents",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(determineTestRunnersCmd).Standalone()

	determineTestRunnersCmd.Flags().Bool("all-supported", false, "Instead of selecting runners based on the chosen formula, return all supported runners.")
	determineTestRunnersCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	determineTestRunnersCmd.Flags().Bool("dependents", false, "Determine runners for testing dependents. Requires `--eval-all` or `HOMEBREW_EVAL_ALL`.")
	determineTestRunnersCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae, whether installed or not, to determine testing dependents.")
	determineTestRunnersCmd.Flags().BoolP("help", "h", false, "Show this message.")
	determineTestRunnersCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	determineTestRunnersCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(determineTestRunnersCmd)
}
