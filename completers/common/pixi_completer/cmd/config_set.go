package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_setCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_setCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	config_setCmd.Flags().StringP("path", "p", "", "Path to a local configuration file")
	config_setCmd.Flags().Bool("shared", false, "Operation on the user configuration shared with other rattler-based tools, which only accepts the keys every such tool understands")
	config_setCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	config_setCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	config_setCmd.MarkFlagsMutuallyExclusive("local", "global", "system", "shared", "path")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"path":          carapace.ActionFiles(),
	})
}
