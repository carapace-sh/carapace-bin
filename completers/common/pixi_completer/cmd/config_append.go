package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_appendCmd).Standalone()

	config_appendCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_appendCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_appendCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	config_appendCmd.Flags().StringP("path", "p", "", "Path to a local configuration file")
	config_appendCmd.Flags().Bool("shared", false, "Operation on the user configuration shared with other rattler-based tools, which only accepts the keys every such tool understands")
	config_appendCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	config_appendCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	config_appendCmd.MarkFlagsMutuallyExclusive("local", "global", "system", "shared", "path")
	configCmd.AddCommand(config_appendCmd)

	carapace.Gen(config_appendCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"path":          carapace.ActionFiles(),
	})
}
