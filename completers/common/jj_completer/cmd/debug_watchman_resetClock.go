package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_watchman_resetClockCmd = &cobra.Command{
	Use:   "reset-clock",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_watchman_resetClockCmd).Standalone()

	debug_watchman_resetClockCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_watchmanCmd.AddCommand(debug_watchman_resetClockCmd)
}
