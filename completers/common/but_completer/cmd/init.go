package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a GitButler project from a git repository in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	initCmd.Flags().BoolP("repo", "r", false, "Also initializes a git repository in the current directory if one does not exist")
	rootCmd.AddCommand(initCmd)
}
