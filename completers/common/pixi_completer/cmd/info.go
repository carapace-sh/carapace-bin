package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information about the system, workspace and environments for the current machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")

	infoCmd.Flags().Bool("extended", false, "Show cache and environment size")
	infoCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	infoCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	infoCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	infoCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
