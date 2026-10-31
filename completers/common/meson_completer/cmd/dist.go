package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var distCmd = &cobra.Command{
	Use:   "dist",
	Short: "Generate release archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distCmd).Standalone()

	distCmd.Flags().StringS("C", "C", "", "directory to cd into before running")
	distCmd.Flags().Bool("allow-dirty", false, "Allow even when repository contains uncommitted changes.")
	distCmd.Flags().String("formats", "", "Comma separated list of archive types to create.")
	distCmd.Flags().Bool("include-subprojects", false, "Include source code of subprojects that have been used for the build.")
	distCmd.Flags().Bool("no-tests", false, "Do not build and test generated packages.")
	distCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (e.g. for compilation and testing).")
	applyCommonFlagCompletion(distCmd)
	rootCmd.AddCommand(distCmd)
}
