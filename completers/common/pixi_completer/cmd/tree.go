package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:     "tree",
	Short:   "Show a tree of workspace dependencies",
	Aliases: []string{"t"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(treeCmd).Standalone()

	treeCmd.Flags().StringP("environment", "e", "", "The environment to list packages for. Defaults to the default environment")
	treeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	treeCmd.Flags().BoolP("invert", "i", false, "Invert tree and show what depends on given package in the regex argument")
	treeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	treeCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	treeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	treeCmd.Flags().StringP("platform", "p", "", "The platform to list packages for. Defaults to the current platform")
	rootCmd.AddCommand(treeCmd)

	carapace.Gen(treeCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
