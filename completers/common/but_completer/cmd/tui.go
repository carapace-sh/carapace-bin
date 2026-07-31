package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:     "tui",
	Short:   "Show an interactive TUI",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "other commands",
}

func init() {
	carapace.Gen(tuiCmd).Standalone()

	tuiCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tuiCmd.Flags().Bool("remember-selection", false, "When the TUI quits save the selection and restore it when re-opening")
	rootCmd.AddCommand(tuiCmd)
}
