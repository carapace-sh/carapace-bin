package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Manage your dotfiles with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfilesCmd).Standalone()

	dotfilesCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(dotfilesCmd)
}
