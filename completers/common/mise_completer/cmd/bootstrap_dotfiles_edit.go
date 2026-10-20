// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a dotfile's source in your editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_editCmd).Standalone()

	bootstrap_dotfiles_editCmd.Flags().Bool("apply", false, "Apply this target after the editor exits")
	bootstrap_dotfiles_editCmd.Flags().String("group", "", "Dotfile group whose tree holds the target, when several groups deploy into its directory")
	bootstrap_dotfiles_editCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_editCmd.Flags().StringP("mode", "m", "", "Mode to use if the target is not managed yet: symlink, symlink-each, copy, or template (default: the `dotfiles.default_mode` setting)")
	bootstrap_dotfiles_editCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	bootstrap_dotfiles_editCmd.Flags().StringP("source", "s", "", "Source path to use if the target is not yet managed")
	bootstrap_dotfiles_editCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt when adding an unmanaged target")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_editCmd)
}
