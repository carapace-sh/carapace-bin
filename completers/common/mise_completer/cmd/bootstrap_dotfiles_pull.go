// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Write fetched changes from other machines to tracked files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_pullCmd).Standalone()

	bootstrap_dotfiles_pullCmd.Flags().BoolP("dry-run", "n", false, "Show the plan without changing anything")
	bootstrap_dotfiles_pullCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_pullCmd.Flags().StringSlice("keep-local", nil, "Resolve a conflict by keeping this machine's version; the next sync pushes it")
	bootstrap_dotfiles_pullCmd.Flags().Bool("keep-local-all", false, "Resolve every remaining conflict by keeping this machine's version")
	bootstrap_dotfiles_pullCmd.Flags().StringSlice("take-remote", nil, "Resolve a conflict with the origin's version")
	bootstrap_dotfiles_pullCmd.Flags().Bool("take-remote-all", false, "Resolve every remaining conflict with the origin's version")
	bootstrap_dotfiles_pullCmd.Flags().BoolP("yes", "y", false, "Accepted for compatibility; pull skips the apply confirmation")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_pullCmd)
}
