package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_title_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the outer terminal title",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_title_clearCmd).Standalone()

	terminal_titleCmd.AddCommand(terminal_title_clearCmd)
}
