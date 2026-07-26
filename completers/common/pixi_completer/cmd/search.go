package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a conda package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringSliceP("channel", "c", nil, "The channels to consider as a name or a url. Multiple channels can be specified by using this field multiple times")
	searchCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	searchCmd.Flags().Bool("json", false, "Output in JSON format")
	searchCmd.Flags().StringP("limit", "l", "5", "Limit the number of versions shown per package, -1 for no limit")
	searchCmd.Flags().StringP("limit-packages", "n", "5", "Limit the number of packages shown, -1 for no limit")
	searchCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	searchCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	searchCmd.Flags().StringP("platform", "p", "", "The platform to search packages for. By default, searches all platforms from the manifest (or all known platforms if no manifest is found). Accepts a workspace platform name; a bare conda subdir (e.g. `linux-64`) is also accepted")
	searchCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"config-file":   carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
