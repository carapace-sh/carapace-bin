// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_history_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a checkpoint: its trigger, changes, and operation journal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_history_showCmd).Standalone()

	bootstrap_dotfiles_history_showCmd.Flags().Bool("files", false, "List every file in the snapshot")
	bootstrap_dotfiles_history_showCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_history_showCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_dotfiles_history_showCmd.Flags().String("path", "", "Resolve `latest~N` among the checkpoints where this path changed")
	bootstrap_dotfiles_historyCmd.AddCommand(bootstrap_dotfiles_history_showCmd)
}
