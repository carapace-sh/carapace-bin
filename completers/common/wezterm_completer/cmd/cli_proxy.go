package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cli_proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "start rpc proxy pipe",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_proxyCmd).Standalone()

	cli_proxyCmd.Flags().BoolP("help", "h", false, "Print help")
	cliCmd.AddCommand(cli_proxyCmd)
}
