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

	lockCmd.Flags().Bool("check", false, "Check if the lockfile is up-to-date")
	lockCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	lockCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	lockCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	rootCmd.AddCommand(lockCmd)

	carapace.Gen(lockCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
