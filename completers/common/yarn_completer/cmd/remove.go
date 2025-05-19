package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove dependencies from the project",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("all", "A", false, "Apply the operation to all workspaces from the current project")
	removeCmd.Flags().String("mode", "", "Change what artifacts installs generate")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionStyledValuesDescribed(
			"skip-build", "do not run the build scripts at all",
			"update-lockfile", "skip the link step altogether",
		),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		yarn.ActionDependencies(),
	)
}
