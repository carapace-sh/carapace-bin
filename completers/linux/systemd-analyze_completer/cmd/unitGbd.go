package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unitGbdCmd = &cobra.Command{
	Use:   "unit-gdb",
	Short: "Attach a debugger to the given running service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unitGbdCmd).Standalone()

	rootCmd.AddCommand(unitGbdCmd)

	carapace.Gen(unitGbdCmd).PositionalCompletion(
		action.ActionServices(unitGbdCmd),
	)
}
