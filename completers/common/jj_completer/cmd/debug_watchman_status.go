package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_watchman_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check whether `watchman` is enabled and whether it's correctly installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_watchman_statusCmd).Standalone()

	debug_watchman_statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_watchmanCmd.AddCommand(debug_watchman_statusCmd)
}
