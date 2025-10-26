package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsp_tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Run the language server and listen on a TCP address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsp_tcpCmd).Standalone()

	lsp_tcpCmd.Flags().String("address", "", "The address to listen on")
	lspCmd.AddCommand(lsp_tcpCmd)
}
