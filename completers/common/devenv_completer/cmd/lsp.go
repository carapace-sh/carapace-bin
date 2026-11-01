package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{
	Use:   "lsp",
	Short: "Start the nixd language server for devenv.nix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lspCmd).Standalone()

	lspCmd.Flags().Bool("print-config", false, "Print nixd configuration and exit")

	rootCmd.AddCommand(lspCmd)
}
