package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netioCmd = &cobra.Command{
	Use:   "netio",
	Short: "Network I/O configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netioCmd).Standalone()
	rootCmd.AddCommand(netioCmd)
}
