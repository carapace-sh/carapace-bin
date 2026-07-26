package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Solve environment and update the lock file without installing the environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()
	lockCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")

	lockCmd.Flags().Bool("check", false, "Check if any changes have been made to the lock file. If yes, exit with a non-zero code")
	lockCmd.Flags().Bool("dry-run", false, "Compute the lock file without writing to disk. Implies --no-install")
	lockCmd.Flags().Bool("json", false, "Output the changes in JSON format")
	lockCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	lockCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	lockCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	lockCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(lockCmd)

	carapace.Gen(lockCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
