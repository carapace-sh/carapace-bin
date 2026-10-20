package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Browse saved versions of your tracked files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_historyCmd).Standalone()

	dotfiles_historyCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_historyCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	dotfiles_historyCmd.Flags().String("label", "", "Only checkpoints with this label")
	dotfiles_historyCmd.Flags().StringP("limit", "n", "20", "Show at most this many checkpoints (0 for all)")
	dotfiles_historyCmd.Flags().String("path", "", "Only checkpoints where this path (or something under it) changed")
	dotfiles_historyCmd.Flags().Bool("pending", false, "Only checkpoints recorded by operations that did not finish")
	dotfiles_historyCmd.Flags().String("trigger", "", "Only checkpoints with this trigger, such as edit, save, agent, bootstrap, apply, rollback, or undo")
	dotfilesCmd.AddCommand(dotfiles_historyCmd)
}
