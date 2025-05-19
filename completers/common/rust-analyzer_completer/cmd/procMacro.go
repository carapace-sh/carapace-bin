package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var procMacroCmd = &cobra.Command{
	Use:   "proc-macro",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(procMacroCmd).Standalone()

	rootCmd.AddCommand(procMacroCmd)
}
