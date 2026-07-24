package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the packages of the current workspace",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")

	listCmd.Flags().StringP("environment", "e", "", "The environment to list packages for. Defaults to the default environment")
	listCmd.Flags().BoolP("explicit", "x", false, "Only list packages that are explicitly defined in the workspace")
	listCmd.Flags().StringSlice("fields", []string{"name", "version", "build", "size", "kind", "source"}, "Select which fields to display and in what order (comma-separated)")
	listCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lock file, doesn't update lock file if it isn't up-to-date with the manifest file")
	listCmd.Flags().Bool("json", false, "Whether to output in json format")
	listCmd.Flags().Bool("json-pretty", false, "Whether to output in json format")
	listCmd.Flags().Bool("locked", false, "Check if lock file is up-to-date before installing the environment, aborts when lock file isn't up-to-date with the manifest file")
	listCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	listCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	listCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	listCmd.Flags().String("platform", "", "The platform to list packages for. Defaults to the platform best matching this machine. Accepts a workspace platform name; a bare conda subdir (e.g. `linux-64`) is also accepted")
	listCmd.Flags().String("sort-by", "name", "Sorting strategy")
	listCmd.Flag("json-pretty").Hidden = true
	listCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
		"sort-by":       carapace.ActionValues("size", "name", "kind"),
	})
}
