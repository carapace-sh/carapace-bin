package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List configuration values",
	Aliases: []string{"ls", "l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	config_listCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths")
	config_listCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_listCmd.Flags().Bool("json", false, "Output in JSON format")
	config_listCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_listCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	config_listCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files")
	config_listCmd.Flags().StringP("path", "p", "", "Path to a local configuration file")
	config_listCmd.Flags().Bool("shared", false, "Operation on the user configuration shared with other rattler-based tools, which only accepts the keys every such tool understands")
	config_listCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	config_listCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	config_listCmd.MarkFlagsMutuallyExclusive("local", "global", "system", "shared", "path")
	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).FlagCompletion(carapace.ActionMap{
		"config-file":   carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"path":          carapace.ActionFiles(),
	})
}
