package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Reverse the tracked-file changes from an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_undoCmd).Standalone()

	dotfiles_undoCmd.Flags().BoolP("dry-run", "n", false, "Show the plan without changing anything")
	dotfiles_undoCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_undoCmd.Flags().BoolP("yes", "y", false, "Apply without prompting")
	dotfilesCmd.AddCommand(dotfiles_undoCmd)
}
