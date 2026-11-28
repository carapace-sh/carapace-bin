package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var gdbtuiCmd = &cobra.Command{
	Use:   "gdbtui",
	Short: "GDB TUI mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gdbtuiCmd).Standalone()
	common.AddGdbFlags(gdbtuiCmd)
	rootCmd.AddCommand(gdbtuiCmd)
}
