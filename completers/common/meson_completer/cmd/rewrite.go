package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewriteCmd = &cobra.Command{
	Use:   "rewrite",
	Short: "Modify the project definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewriteCmd).Standalone()

	rewriteCmd.Flags().BoolP("skip-errors", "S", false, "Skip errors instead of aborting")
	rewriteCmd.Flags().StringP("sourcedir", "s", "", "Path to source directory.")
	rewriteCmd.Flags().BoolP("verbose", "V", false, "Enable verbose output")
	applyCommonFlagCompletion(rewriteCmd)
	rootCmd.AddCommand(rewriteCmd)
}
