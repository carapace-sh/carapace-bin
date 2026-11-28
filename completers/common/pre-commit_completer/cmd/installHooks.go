package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installHooksCmd = &cobra.Command{
	Use:   "install-hooks",
	Short: "Install hook environments for all environments in the config file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installHooksCmd).Standalone()

	installHooksCmd.Flags().String("color", "", "Whether to use color in output.")
	installHooksCmd.Flags().StringP("config", "c", "", "Path to alternate config file")

	carapace.Gen(installHooksCmd).FlagCompletion(carapace.ActionMap{
		"color":  action.ActionColors(),
		"config": carapace.ActionFiles(".yaml", ".yml"),
	})

	rootCmd.AddCommand(installHooksCmd)
}
