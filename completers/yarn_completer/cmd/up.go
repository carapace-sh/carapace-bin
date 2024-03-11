package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:     "up",
	Short:   "Upgrade dependencies across the project",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().BoolP("caret", "C", false, "Use the ^ semver modifier on the resolved range")
	upCmd.Flags().BoolP("exact", "E", false, "Don't use any semver modifier on the resolved range")
	upCmd.Flags().BoolP("interactive", "i", false, "Offer various choices, depending on the detected upgrade paths")
	upCmd.Flags().String("mode", "", "Change what artifacts installs generate")
	upCmd.Flags().BoolP("recursive", "R", false, "Resolve again ALL resolutions for those packages")
	upCmd.Flags().BoolP("tilde", "T", false, "Use the ~ semver modifier on the resolved range")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValuesDescribed(
			"skip-build", "do not run the build scripts at all",
			"update-lockfile", "skip the link step altogether",
		),
	})

	carapace.Gen(upCmd).PositionalAnyCompletion(
		yarn.ActionDependencies(),
	)
}
