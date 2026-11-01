package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the state of your dotfiles and their history",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_statusCmd).Standalone()

	dotfiles_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	dotfiles_statusCmd.Flags().Bool("missing", false, "Exit with code 1 if any dotfile is not in its desired state (missing, source missing, or differs)")
	dotfiles_statusCmd.Flags().Bool("no-truncate", false, "Truncate long terminal output to fit the available width")
	dotfiles_statusCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	dotfiles_statusCmd.Flags().Bool("truncate", true, "Truncate long terminal output to fit the available width")
	dotfiles_statusCmd.Flags().BoolP("verbose", "v", false, "List every entry, including those already applied or tracked")
	dotfiles_statusCmd.Flag("no-truncate").Hidden = true
	dotfilesCmd.AddCommand(dotfiles_statusCmd)
}
