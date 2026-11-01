package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_history_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Set the description of a checkpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_history_describeCmd).Standalone()

	dotfiles_history_describeCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_historyCmd.AddCommand(dotfiles_history_describeCmd)
}
