package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_history_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Compare checkpoints, or your files against one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_history_diffCmd).Standalone()

	bootstrap_dotfiles_history_diffCmd.Flags().Bool("exit-code", false, "Exit 1 when the two sides differ")
	bootstrap_dotfiles_history_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_history_diffCmd.Flags().Bool("operation", false, "Compare an operation with the checkpoint it saved before it ran")
	bootstrap_dotfiles_history_diffCmd.Flags().BoolP("patch", "p", false, "Print the full patch instead of a per-file summary")
	bootstrap_dotfiles_history_diffCmd.Flags().String("path", "", "Restrict to one path (a file or a directory)")
	bootstrap_dotfiles_historyCmd.AddCommand(bootstrap_dotfiles_history_diffCmd)
}
