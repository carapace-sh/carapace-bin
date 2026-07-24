package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleanup the environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("activation-cache", false, "Only remove the activation cache")
	cleanCmd.Flags().Bool("build", false, "Only remove the pixi-build cache")
	cleanCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	cleanCmd.Flags().StringP("environment", "e", "", "The environment directory to remove")
	cleanCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	cleanCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	cleanCmd.Flags().Bool("workspaces-registry", false, "Only remove disassociated workspace registries")
	cleanCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
	})
}
