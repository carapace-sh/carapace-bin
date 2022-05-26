package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "build a derivation or fetch a store path",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("json", false, "Produce output in JSON format, suitable for consumption by another program.")
	buildCmd.Flags().Bool("no-link", false, "Do not create symlinks to the build results.")
	buildCmd.Flags().StringP("out-link", "o", "", "Use path as prefix for the symlinks to the build results.")
	buildCmd.Flags().Bool("print-out-paths", false, "Print the resulting output paths")
	buildCmd.Flags().String("profile", "", "The profile to update.")
	buildCmd.Flags().Bool("rebuild", false, "Rebuild an already built package and compare the result to the existing store paths.")

	addEvaluationFlags(buildCmd)
	addFlakeFlags(buildCmd)
	addInterpretationFlags(buildCmd)
	addLoggingFlags(buildCmd)

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"out-link": carapace.ActionFiles(),
		"profile":  carapace.ActionFiles(),
	})
}
