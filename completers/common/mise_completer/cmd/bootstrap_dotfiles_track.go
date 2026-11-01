package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track a file or directory in place",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_trackCmd).Standalone()

	bootstrap_dotfiles_trackCmd.Flags().Bool("allow-plaintext", false, "Save an explicitly tracked credential-named file in plaintext")
	bootstrap_dotfiles_trackCmd.Flags().BoolP("dry-run", "n", false, "Show what each path expands to (files, size, what is left out) without tracking it")
	bootstrap_dotfiles_trackCmd.Flags().Bool("encrypt", false, "Encrypt contents before saving them to history (requires `[history.encryption] recipients`)")
	bootstrap_dotfiles_trackCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_trackCmd.Flags().Bool("local", false, "Keep this file's history on this machine only; it is never shared")
	bootstrap_dotfiles_trackCmd.Flags().Bool("machine", false, "Keep a separate stream on every machine, never applied on another")
	bootstrap_dotfiles_trackCmd.Flags().Bool("no-autosave", false, "Save only on `mise dot save <path>`, never automatically")
	bootstrap_dotfiles_trackCmd.Flags().String("os", "", "Declare a variant for this platform (macos, linux, linux/arm64, …)")
	bootstrap_dotfiles_trackCmd.Flags().String("profile", "", "Declare a variant for this config environment (`MISE_ENV` or `-E`)")
	bootstrap_dotfiles_trackCmd.Flags().BoolP("yes", "y", false, "Accepted for compatibility; track no longer asks to confirm the paths")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_trackCmd)
}
