package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run local TLS proxy allowing connecting to Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxyCmd).Standalone()

	proxyCmd.Flags().String("proxy-log-output", "stderr", "Select where proxy status messages are printed. Defaults to \"stderr\", but can be used to revert to legacy behavior of send proxy output to stdout. Valid values are \"stdout\", \"stderr\", and \"none\".")
	rootCmd.AddCommand(proxyCmd)
}
