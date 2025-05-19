package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add dependencies to the project",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("cached", false, "Reuse the highest version already used somewhere within the project")
	addCmd.Flags().BoolP("caret", "C", false, "Use the ^ semver modifier on the resolved range")
	addCmd.Flags().BoolP("dev", "D", false, "Add a package as a dev dependency")
	addCmd.Flags().BoolP("exact", "E", false, "Don't use any semver modifier on the resolved range")
	addCmd.Flags().BoolP("interactive", "i", false, "Reuse the specified package from other workspaces in the project")
	addCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	addCmd.Flags().String("mode", "", "Change what artifacts installs generate")
	addCmd.Flags().BoolP("optional", "O", false, "Add / upgrade a package to an optional regular / peer dependency")
	addCmd.Flags().BoolP("peer", "P", false, "Add a package as a peer dependency")
	addCmd.Flags().Bool("prefer-dev", false, "Add / upgrade a package to a dev dependency")
	addCmd.Flags().BoolP("tilde", "T", false, "Use the ~ semver modifier on the resolved range")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValuesDescribed(
			"skip-build", "do not run the build scripts at all",
			"update-lockfile", "skip the link step altogether",
		),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		npm.ActionPackageSearch(""),
	)
}
