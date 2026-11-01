package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Browse metadata interactively in a full-screen terminal UI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tuiCmd).Standalone()

	tuiCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(tuiCmd)
}
