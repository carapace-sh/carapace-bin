package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the background daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	daemonCmd.Flags().Bool("daemonize", false, "Internal flag for daemonization")
	daemonCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonCmd.Flags().Bool("show-logs", false, "Also write daemon logs to the console (useful for debugging)")
	daemonCmd.Flag("daemonize").Hidden = true
	rootCmd.AddCommand(daemonCmd)
}
