package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the pre-commit script.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("allow-missing-config", false, "Whether to allow a missing `pre-commit` configuration file or exit with a failure code.")
	installCmd.Flags().String("color", "", "Whether to use color in output.")
	installCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	installCmd.Flags().StringP("hook-type", "t", "", "")
	installCmd.Flags().Bool("install-hooks", false, "Whether to install hook environments for all environments in the config file.")
	installCmd.Flags().BoolP("overwrite", "f", false, "Overwrite existing hooks / remove migration mode.")

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"color":     action.ActionColors(),
		"config":    carapace.ActionFiles(".yaml", ".yml"),
		"hook-type": action.ActionHookTypes(),
	})

	rootCmd.AddCommand(installCmd)
}
