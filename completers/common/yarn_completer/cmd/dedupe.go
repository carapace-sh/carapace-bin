package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dedupeCmd = &cobra.Command{
	Use:     "dedupe",
	Short:   "Deduplicate dependencies with overlapping ranges",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dedupeCmd).Standalone()

	dedupeCmd.Flags().BoolP("check", "c", false, "Exit with exit code 1 when duplicates are found, without persisting the dependency tree")
	dedupeCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	dedupeCmd.Flags().String("mode", "", "Change what artifacts installs generate")
	dedupeCmd.Flags().StringP("strategy", "s", "", "The strategy to use when deduping dependencies")
	rootCmd.AddCommand(dedupeCmd)

	carapace.Gen(dedupeCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValuesDescribed(
			"skip-build", "do not run the build scripts at all",
			"update-lockfile", "skip the link step altogether",
		),
		"strategy": carapace.ActionValuesDescribed(
			"highest", "reuse the locators with the highest versions",
		),
	})

	// TODO positional completion
}
