package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var gdbCmd = &cobra.Command{
	Use:   "gdb",
	Short: "Run the GDB",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gdbCmd).Standalone()

	common.AddGdbFlags(gdbCmd)
	gdbCmd.Flags().Bool("batch", false, "Exit after processing gdbinit")
	gdbCmd.Flags().String("gdb-tui", "", "Run gdb in TUI mode")
	gdbCmd.Flag("batch").Hidden = true
	rootCmd.AddCommand(gdbCmd)
}
