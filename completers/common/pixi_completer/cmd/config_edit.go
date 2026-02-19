package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_editCmd).Standalone()

	config_editCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_editCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_editCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	config_editCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	configCmd.AddCommand(config_editCmd)

	carapace.Gen(config_editCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
