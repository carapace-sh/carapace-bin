package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Stop tracking a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_untrackCmd).Standalone()

	dotfiles_untrackCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfilesCmd.AddCommand(dotfiles_untrackCmd)
}
