package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/typst_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches an input file and recompiles on changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	watchCmd.Flags().Bool("no-reload", false, "Disables the injected live reload script for HTML export.")
	watchCmd.Flags().Bool("no-serve", false, "Disables the built-in HTTP server for HTML export")
	watchCmd.Flags().String("port", "", "The port where HTML is served.")

	common.AddPdfFlags(watchCmd)

	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})

	carapace.Gen(watchCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
