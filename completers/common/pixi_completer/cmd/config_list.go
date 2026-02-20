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

	config_listCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_listCmd.Flags().Bool("json", false, "Output in JSON format")
	config_listCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	config_listCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
