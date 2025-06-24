package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var syscallFilterCmd = &cobra.Command{
	Use:   "syscall-filter",
	Short: "List syscalls in seccomp filters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syscallFilterCmd).Standalone()

	rootCmd.AddCommand(syscallFilterCmd)

	carapace.Gen(syscallFilterCmd).PositionalAnyCompletion(
		action.ActionSyscallSets(),
	)
}
