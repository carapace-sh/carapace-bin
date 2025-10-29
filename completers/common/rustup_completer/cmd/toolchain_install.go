package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var toolchain_installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install or update the given toolchains, or by default the active toolchain",
	Aliases: []string{"update", "add"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_installCmd).Standalone()

	toolchain_installCmd.Flags().Bool("allow-downgrade", false, "Allow rustup to downgrade the toolchain to satisfy your component choice")
	toolchain_installCmd.Flags().StringSliceP("component", "c", nil, "Comma-separated list of components to be added on installation")
	toolchain_installCmd.Flags().Bool("force", false, "Force an update, even if some components are missing")
	toolchain_installCmd.Flags().Bool("force-non-host", false, "Install toolchains that require an emulator. See https://github.com/rust-lang/rustup/wiki/Non-host-toolchains")
	toolchain_installCmd.Flags().BoolP("help", "h", false, "Print help")
	toolchain_installCmd.Flags().Bool("no-self-update", false, "Don't perform self update when running the `rustup toolchain install` command")
	toolchain_installCmd.Flags().String("profile", "", "")
	toolchain_installCmd.Flags().StringSliceP("target", "t", nil, "Comma-separated list of targets to be added on installation")
	toolchainCmd.AddCommand(toolchain_installCmd)

	carapace.Gen(toolchain_installCmd).FlagCompletion(carapace.ActionMap{
		"component": action.ActionAvailableComponents().UniqueList(","),
		"profile":   carapace.ActionValues("minimal", "default", "complete"),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(false).Invoke(c).ToMultiPartsA("-")
		}),
	})

	carapace.Gen(toolchain_installCmd).PositionalAnyCompletion(
		action.ActionDefaultChannels().FilterArgs(),
	)
}
