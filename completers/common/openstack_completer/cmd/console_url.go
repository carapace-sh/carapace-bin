package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var console_urlCmd = &cobra.Command{
	Use:   "url",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(console_urlCmd).Standalone()

	consoleCmd.AddCommand(console_urlCmd)
}
