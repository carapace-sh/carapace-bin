package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{
	Use:   "lsp",
	Short: "Start LSP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lspCmd).Standalone()

	lspCmd.Flags().StringS("goplsLog", "goplsLog", "", "The file to log gopls output, or leave empty to disable logging.")
	lspCmd.Flags().BoolS("goplsRPCTrace", "goplsRPCTrace", false, "Set gopls to log input and output messages")
	lspCmd.Flags().BoolS("help", "help", false, "Print help and exit")
	lspCmd.Flags().StringS("http", "http", "", "Enable http debug server by setting a listen address")
	lspCmd.Flags().StringS("log", "log", "", "The file to log templ LSP output to, or leave empty to disable logging")
	lspCmd.Flags().BoolS("pprof", "pprof", false, "Enable pprof web server")
	rootCmd.AddCommand(lspCmd)

	carapace.Gen(lspCmd).FlagCompletion(carapace.ActionMap{
		"goplsLog": carapace.ActionFiles(),
		"http": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return net.ActionPorts()
			}
		}),
		"log": carapace.ActionFiles(),
		"pprof": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return net.ActionPorts()
			}
		}),
	})
}
