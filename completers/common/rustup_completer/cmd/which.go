package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Display which binary will be run for a given command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()

	whichCmd.Flags().BoolP("help", "h", false, "Print help")
	whichCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', '1.8.0', or a custom toolchain name. For more information see `rustup help toolchain`")
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})
}
