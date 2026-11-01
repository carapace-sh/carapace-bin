package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover an interrupted dotfile operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_recoverCmd).Standalone()

	bootstrap_dotfiles_recoverCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_recoverCmd.Flags().Bool("keep-current", false, "Keep the files on disk instead of restoring the recovery copies")
	bootstrap_dotfiles_recoverCmd.Flags().BoolP("yes", "y", false, "Discard the operation's recovery copies without prompting")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_recoverCmd)
}
