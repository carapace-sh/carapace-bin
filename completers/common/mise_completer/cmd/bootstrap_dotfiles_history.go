// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Browse saved versions of your tracked files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_historyCmd).Standalone()

	bootstrap_dotfiles_historyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_historyCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_dotfiles_historyCmd.Flags().String("label", "", "Only checkpoints with this label")
	bootstrap_dotfiles_historyCmd.Flags().StringP("limit", "n", "20", "Show at most this many checkpoints (0 for all)")
	bootstrap_dotfiles_historyCmd.Flags().String("path", "", "Only checkpoints where this path (or something under it) changed")
	bootstrap_dotfiles_historyCmd.Flags().Bool("pending", false, "Only checkpoints recorded by operations that did not finish")
	bootstrap_dotfiles_historyCmd.Flags().String("trigger", "", "Only checkpoints with this trigger, such as edit, save, agent, bootstrap, apply, rollback, or undo")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_historyCmd)
}
