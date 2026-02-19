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
	cleanCmd.Flags().StringP("environment", "e", "", "The environment directory to remove")
	cleanCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
	})
}
