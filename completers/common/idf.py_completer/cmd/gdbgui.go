package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var gdbguiCmd = &cobra.Command{
	Use:   "gdbgui",
	Short: "GDB UI in default browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gdbguiCmd).Standalone()

	common.AddGdbFlags(gdbguiCmd)
	gdbguiCmd.Flags().String("gdbgui-port", "", "Port on which gdbgui will be hosted")
	rootCmd.AddCommand(gdbguiCmd)
}
