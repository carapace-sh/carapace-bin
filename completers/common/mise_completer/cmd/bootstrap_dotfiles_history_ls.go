package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_history_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List checkpoints, newest first",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_history_lsCmd).Standalone()

	bootstrap_dotfiles_history_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_history_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_dotfiles_history_lsCmd.Flags().String("label", "", "Only checkpoints with this label")
	bootstrap_dotfiles_history_lsCmd.Flags().StringP("limit", "n", "20", "Show at most this many checkpoints (0 for all)")
	bootstrap_dotfiles_history_lsCmd.Flags().String("path", "", "Only checkpoints where this path (or something under it) changed")
	bootstrap_dotfiles_history_lsCmd.Flags().Bool("pending", false, "Only checkpoints recorded by operations that did not finish")
	bootstrap_dotfiles_history_lsCmd.Flags().String("trigger", "", "Only checkpoints with this trigger, such as edit, save, agent, bootstrap, apply, rollback, or undo")
	bootstrap_dotfiles_historyCmd.AddCommand(bootstrap_dotfiles_history_lsCmd)
}
