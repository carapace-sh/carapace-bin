package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Push saved checkpoints and fetch incoming changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_syncCmd).Standalone()

	bootstrap_dotfiles_syncCmd.Flags().Bool("allow-plaintext-history", false, "Allow pushing older unencrypted versions of files that are now encrypted, and versions that look like they contain secrets")
	bootstrap_dotfiles_syncCmd.Flags().Bool("best-effort", false, "Warn instead of failing when the origin is unreachable")
	bootstrap_dotfiles_syncCmd.Flags().Bool("fetch-only", false, "Fetch without pushing")
	bootstrap_dotfiles_syncCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_syncCmd)
}
