package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_title_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the outer terminal title",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_title_setCmd).Standalone()

	terminal_titleCmd.AddCommand(terminal_title_setCmd)
}
