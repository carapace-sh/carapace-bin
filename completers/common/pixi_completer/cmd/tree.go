package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show a tree of workspace dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(treeCmd).Standalone()

	treeCmd.Flags().StringP("environment", "e", "", "The environment to show the tree for")
	treeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	treeCmd.Flags().BoolP("invert", "i", false, "Invert the dependency tree")
	treeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	treeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	treeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	treeCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	treeCmd.Flags().StringP("platform", "p", "", "The platform to show the tree for")
	rootCmd.AddCommand(treeCmd)

	carapace.Gen(treeCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
