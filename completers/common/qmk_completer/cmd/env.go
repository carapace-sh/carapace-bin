package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Prints environment information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(envCmd)

	carapace.Gen(envCmd).PositionalCompletion(
		action.ActionEnvs(),
	)
}
