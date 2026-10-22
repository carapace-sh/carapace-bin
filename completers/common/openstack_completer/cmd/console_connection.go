package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var console_connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(console_connectionCmd).Standalone()

	consoleCmd.AddCommand(console_connectionCmd)
}
