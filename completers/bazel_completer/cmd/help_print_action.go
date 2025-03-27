package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_print_actionCmd = &cobra.Command{
	Use:   "print_action",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_print_actionCmd).Standalone()

	helpCmd.AddCommand(help_print_actionCmd)
}
