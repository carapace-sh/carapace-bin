package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initTemplatedirCmd = &cobra.Command{
	Use:   "init-templatedir [directory]",
	Short: "Install hook script in a directory intended for use with `git config init.templateDir`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initTemplatedirCmd).Standalone()

	initTemplatedirCmd.Flags().String("color", "", "Whether to use color in output.")
	initTemplatedirCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	initTemplatedirCmd.Flags().StringP("hook-type", "t", "", "")
	initTemplatedirCmd.Flags().Bool("no-allow-missing-config", false, "Assume cloned repos should have a `pre-commit` config.")

	carapace.Gen(initTemplatedirCmd).FlagCompletion(carapace.ActionMap{
		"color":     action.ActionColors(),
		"config":    carapace.ActionFiles(".yaml", ".yml"),
		"hook-type": action.ActionHookTypes(),
	})

	carapace.Gen(initTemplatedirCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	rootCmd.AddCommand(initTemplatedirCmd)
}
