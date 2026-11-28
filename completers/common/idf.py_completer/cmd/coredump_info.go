package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var coredumpInfoCmd = &cobra.Command{
	Use:   "coredump-info",
	Short: "Print crashed task's registers, callstack, and memory regions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coredumpInfoCmd).Standalone()

	common.AddCoredumpFlags(coredumpInfoCmd)
	coredumpInfoCmd.Flags().String("gdb-timeout-sec", "", "Overwrite the default internal delay for gdb responses")
	coredumpInfoCmd.Flags().Bool("print-all-vars", false, "Print all variables tagged with COREDUMP_*_ATTR")
	coredumpInfoCmd.Flags().StringArray("print-var", nil, "Print a variable or expression")
	rootCmd.AddCommand(coredumpInfoCmd)
}
