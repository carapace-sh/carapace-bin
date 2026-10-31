package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hookImplCmd = &cobra.Command{
	Use:    "hook-impl",
	Short:  "Invoke hooks from an installed git hook script",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookImplCmd).Standalone()

	hookImplCmd.Flags().String("color", "", "Whether to use color in output.")
	hookImplCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	hookImplCmd.Flags().String("hook-dir", "", "")
	hookImplCmd.Flags().String("hook-type", "", "")
	hookImplCmd.Flags().Bool("skip-on-missing-config", false, "")
	hookImplCmd.MarkFlagRequired("hook-type")

	carapace.Gen(hookImplCmd).FlagCompletion(carapace.ActionMap{
		"color":     action.ActionColors(),
		"config":    carapace.ActionFiles(".yaml", ".yml"),
		"hook-dir":  carapace.ActionDirectories(),
		"hook-type": action.ActionHookTypes(),
	})

	rootCmd.AddCommand(hookImplCmd)
}
