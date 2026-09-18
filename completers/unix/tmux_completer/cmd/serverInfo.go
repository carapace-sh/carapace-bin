package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverInfoCmd = &cobra.Command{
	Use:     "server-info",
	Aliases: []string{"info"},
	Short:   "show server information",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverInfoCmd).Standalone()

	rootCmd.AddCommand(serverInfoCmd)
}
