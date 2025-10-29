package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var override_unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "Remove the override toolchain for a directory",
	Aliases: []string{"remove", "rm", "delete", "del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_unsetCmd).Standalone()

	override_unsetCmd.Flags().BoolP("help", "h", false, "Print help")
	override_unsetCmd.Flags().Bool("nonexistent", false, "Remove override toolchain for all nonexistent directories")
	override_unsetCmd.Flags().String("path", "", "Path to the directory")
	overrideCmd.AddCommand(override_unsetCmd)

	carapace.Gen(override_unsetCmd).FlagCompletion(carapace.ActionMap{
		"path": action.ActionOverrides(),
	})
}
