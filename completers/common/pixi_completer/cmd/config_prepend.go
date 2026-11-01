package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "Prepend a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_prependCmd).Standalone()

	config_prependCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_prependCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_prependCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	config_prependCmd.Flags().StringP("path", "p", "", "Path to a local configuration file")
	config_prependCmd.Flags().Bool("shared", false, "Operation on the user configuration shared with other rattler-based tools, which only accepts the keys every such tool understands")
	config_prependCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	config_prependCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	config_prependCmd.MarkFlagsMutuallyExclusive("local", "global", "system", "shared", "path")
	configCmd.AddCommand(config_prependCmd)

	carapace.Gen(config_prependCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"path":          carapace.ActionFiles(),
	})
}
