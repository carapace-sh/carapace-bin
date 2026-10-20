// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_history_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Compare checkpoints, or your files against one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_history_diffCmd).Standalone()

	dotfiles_history_diffCmd.Flags().Bool("exit-code", false, "Exit 1 when the two sides differ")
	dotfiles_history_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_history_diffCmd.Flags().Bool("operation", false, "Compare an operation with the checkpoint it saved before it ran")
	dotfiles_history_diffCmd.Flags().BoolP("patch", "p", false, "Print the full patch instead of a per-file summary")
	dotfiles_history_diffCmd.Flags().String("path", "", "Restrict to one path (a file or a directory)")
	dotfiles_historyCmd.AddCommand(dotfiles_history_diffCmd)
}
