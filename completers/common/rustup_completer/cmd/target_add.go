package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var target_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a target to a Rust toolchain",
	Aliases: []string{"install"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_addCmd).Standalone()

	target_addCmd.Flags().BoolP("help", "h", false, "Print help")
	target_addCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see `rustup help toolchain`")
	targetCmd.AddCommand(target_addCmd)

	carapace.Gen(target_addCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})

	carapace.Gen(target_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(false).Invoke(c).Filter(c.Args...).ToMultiPartsA("-") // TODO user FilterArgs
		}),
	)
}
