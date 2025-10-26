package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsp_stdioCmd = &cobra.Command{
	Use:   "stdio",
	Short: "Run the language server over the standard input and output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsp_stdioCmd).Standalone()

	lspCmd.AddCommand(lsp_stdioCmd)
}
