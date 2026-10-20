package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Save tracked files as they change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_watchCmd).Standalone()

	bootstrap_dotfiles_watchCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_watchCmd.Flags().BoolP("json", "J", false, "Print one JSON object per line instead of log lines")
	bootstrap_dotfiles_watchCmd.Flags().Bool("once", false, "Rescan and sync once, then exit (for timers and cron)")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_watchCmd)
}
