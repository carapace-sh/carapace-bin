package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tcpipCmd = &cobra.Command{
	Use:   "tcpip",
	Short: "restart adbd listening on TCP on PORT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcpipCmd).Standalone()

	rootCmd.AddCommand(tcpipCmd)
}
