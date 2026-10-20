package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_history_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a checkpoint: its trigger, changes, and operation journal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_history_showCmd).Standalone()

	dotfiles_history_showCmd.Flags().Bool("files", false, "List every file in the snapshot")
	dotfiles_history_showCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_history_showCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	dotfiles_history_showCmd.Flags().String("path", "", "Resolve `latest~N` among the checkpoints where this path changed")
	dotfiles_historyCmd.AddCommand(dotfiles_history_showCmd)
}
