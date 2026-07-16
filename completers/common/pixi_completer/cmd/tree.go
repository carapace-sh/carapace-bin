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
	treeCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")

	treeCmd.Flags().StringP("environment", "e", "", "The environment to list packages for. Defaults to the default environment")
	treeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lock file, doesn't update lock file if it isn't up-to-date with the manifest file")
	treeCmd.Flags().BoolP("invert", "i", false, "Invert tree and show what depends on given package in the regex argument")
	treeCmd.Flags().Bool("locked", false, "Check if lock file is up-to-date before installing the environment, aborts when lock file isn't up-to-date with the manifest file")
	treeCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	treeCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	treeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	treeCmd.Flags().StringP("platform", "p", "", "The platform to list packages for. Defaults to the platform best matching this machine. Accepts a workspace platform name; a bare conda subdir (e.g. `linux-64`) is also accepted")
	treeCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(treeCmd)

	carapace.Gen(treeCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
