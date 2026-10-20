// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Save tracked files as they change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_watchCmd).Standalone()

	dotfiles_watchCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_watchCmd.Flags().BoolP("json", "J", false, "Print one JSON object per line instead of log lines")
	dotfiles_watchCmd.Flags().Bool("once", false, "Rescan and sync once, then exit (for timers and cron)")
	dotfilesCmd.AddCommand(dotfiles_watchCmd)
}
