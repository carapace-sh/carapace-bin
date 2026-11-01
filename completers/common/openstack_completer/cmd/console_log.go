package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var console_logCmd = &cobra.Command{
	Use:   "log",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(console_logCmd).Standalone()

	consoleCmd.AddCommand(console_logCmd)
}
