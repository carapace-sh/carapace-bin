package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use:     "gui",
	Short:   "Open the GitButler GUI for the current project.",
	Aliases: []string{"."},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guiCmd).Standalone()

	guiCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(guiCmd)
}
