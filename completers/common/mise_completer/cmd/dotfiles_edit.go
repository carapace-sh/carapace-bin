package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a dotfile's source in your editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_editCmd).Standalone()

	dotfiles_editCmd.Flags().Bool("apply", false, "Apply this target after the editor exits")
	dotfiles_editCmd.Flags().String("group", "", "Dotfile group whose tree holds the target, when several groups deploy into its directory")
	dotfiles_editCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_editCmd.Flags().StringP("mode", "m", "", "Mode to use if the target is not managed yet: symlink, symlink-each, copy, or template (default: the `dotfiles.default_mode` setting)")
	dotfiles_editCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	dotfiles_editCmd.Flags().StringP("source", "s", "", "Source path to use if the target is not yet managed")
	dotfiles_editCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt when adding an unmanaged target")
	dotfilesCmd.AddCommand(dotfiles_editCmd)
}
