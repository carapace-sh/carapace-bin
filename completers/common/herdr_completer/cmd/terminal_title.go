package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_titleCmd = &cobra.Command{
	Use:   "title",
	Short: "Manage the outer terminal title",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_titleCmd).Standalone()

	terminalCmd.AddCommand(terminal_titleCmd)
}
