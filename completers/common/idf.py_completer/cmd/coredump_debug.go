package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var coredumpDebugCmd = &cobra.Command{
	Use:   "coredump-debug",
	Short: "Create core dump ELF file and run GDB debug session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coredumpDebugCmd).Standalone()
	common.AddCoredumpFlags(coredumpDebugCmd)
	rootCmd.AddCommand(coredumpDebugCmd)
}
