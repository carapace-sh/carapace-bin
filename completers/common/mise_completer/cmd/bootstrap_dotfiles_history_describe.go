package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_history_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Set the description of a checkpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_history_describeCmd).Standalone()

	bootstrap_dotfiles_history_describeCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_historyCmd.AddCommand(bootstrap_dotfiles_history_describeCmd)
}
