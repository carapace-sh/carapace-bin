package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_unsetCmd).Standalone()

	config_unsetCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_unsetCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_unsetCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	config_unsetCmd.Flags().StringP("path", "p", "", "Path to a local configuration file")
	config_unsetCmd.Flags().Bool("shared", false, "Operation on the user configuration shared with other rattler-based tools, which only accepts the keys every such tool understands")
	config_unsetCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	config_unsetCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	config_unsetCmd.MarkFlagsMutuallyExclusive("local", "global", "system", "shared", "path")
	configCmd.AddCommand(config_unsetCmd)

	carapace.Gen(config_unsetCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"path":          carapace.ActionFiles(),
	})
}
