package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the pre-commit script.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().String("color", "", "Whether to use color in output.")
	uninstallCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	uninstallCmd.Flags().StringP("hook-type", "t", "", "")

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"color":     action.ActionColors(),
		"config":    carapace.ActionFiles(".yaml", ".yml"),
		"hook-type": action.ActionHookTypes(),
	})

	rootCmd.AddCommand(uninstallCmd)
}
