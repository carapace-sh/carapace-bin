package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the state of your dotfiles and their history",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_statusCmd).Standalone()

	bootstrap_dotfiles_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_dotfiles_statusCmd.Flags().Bool("missing", false, "Exit with code 1 if any dotfile is not in its desired state (missing, source missing, or differs)")
	bootstrap_dotfiles_statusCmd.Flags().Bool("no-truncate", false, "Truncate long terminal output to fit the available width")
	bootstrap_dotfiles_statusCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	bootstrap_dotfiles_statusCmd.Flags().Bool("truncate", true, "Truncate long terminal output to fit the available width")
	bootstrap_dotfiles_statusCmd.Flags().BoolP("verbose", "v", false, "List every entry, including those already applied or tracked")
	bootstrap_dotfiles_statusCmd.Flag("no-truncate").Hidden = true
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_statusCmd)
}
