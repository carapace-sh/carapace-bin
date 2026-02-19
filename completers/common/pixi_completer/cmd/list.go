package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the packages of the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringP("environment", "e", "", "The environment to list packages for")
	listCmd.Flags().BoolP("explicit", "x", false, "Only list packages that are explicitly defined")
	listCmd.Flags().String("fields", "", "Fields to show")
	listCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	listCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	listCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	listCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	listCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	listCmd.Flags().String("platform", "", "The platform for which to list packages")
	listCmd.Flags().String("sort-by", "", "Sorting strategy")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
		"sort-by":       carapace.ActionValues("size", "name", "kind"),
	})
}
