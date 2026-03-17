package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Terminal emulator for atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hexCmd).Standalone()

	helpCmd.AddCommand(help_hexCmd)
}
