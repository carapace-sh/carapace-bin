package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfilesCmd = &cobra.Command{
	Use:     "dotfiles",
	Short:   "Deploy, track, and sync dotfiles",
	Aliases: []string{"dot"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfilesCmd).Standalone()

	dotfilesCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfilesCmd.Flags().Bool("local", false, "Work on this machine's local-only history (`mode = \"track-local\"`), which is never shared")
	rootCmd.AddCommand(dotfilesCmd)
}
