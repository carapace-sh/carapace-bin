package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update Rust toolchains and rustup",
	Aliases: []string{"upgrade", "up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("force", false, "Force an update, even if some components are missing")
	updateCmd.Flags().Bool("force-non-host", false, "Install toolchains that require an emulator. See https://github.com/rust-lang/rustup/wiki/Non-host-toolchains")
	updateCmd.Flags().BoolP("help", "h", false, "Print help")
	updateCmd.Flags().Bool("no-self-update", false, "Don't perform self update when running the `rustup update` command")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		action.ActionToolchains().FilterArgs(),
	)
}
