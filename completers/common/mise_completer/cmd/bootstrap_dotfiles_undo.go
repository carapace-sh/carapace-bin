package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Reverse the tracked-file changes from an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_undoCmd).Standalone()

	bootstrap_dotfiles_undoCmd.Flags().BoolP("dry-run", "n", false, "Show the plan without changing anything")
	bootstrap_dotfiles_undoCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_undoCmd.Flags().BoolP("yes", "y", false, "Apply without prompting")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_undoCmd)
}
