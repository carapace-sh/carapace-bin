package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_watchman_queryChangedFilesCmd = &cobra.Command{
	Use:   "query-changed-files",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_watchman_queryChangedFilesCmd).Standalone()

	debug_watchman_queryChangedFilesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_watchmanCmd.AddCommand(debug_watchman_queryChangedFilesCmd)
}