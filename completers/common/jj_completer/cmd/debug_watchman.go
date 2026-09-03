package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_watchmanCmd = &cobra.Command{
	Use:   "watchman",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_watchmanCmd).Standalone()

	debug_watchmanCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_watchmanCmd)
}

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

var debug_watchman_queryClockCmd = &cobra.Command{
	Use:   "query-clock",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_watchman_queryClockCmd).Standalone()

	debug_watchman_queryClockCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_watchmanCmd.AddCommand(debug_watchman_queryClockCmd)
}

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

var debug_workingCopyCmd = &cobra.Command{
	Use:   "working-copy",
	Short: "Show information about the working copy state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_workingCopyCmd).Standalone()

	debug_workingCopyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_workingCopyCmd)
}