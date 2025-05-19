package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Set the default toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultCmd).Standalone()

	defaultCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(defaultCmd)

	carapace.Gen(defaultCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionDefaultChannels(),
			action.ActionToolchains(),
		).ToA(),
	)
}
