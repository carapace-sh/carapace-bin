package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a GitButler project from a git repository in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	setupCmd.Flags().Bool("init", false, "Initialize a new git repository with an empty commit if one doesn't exist")
	rootCmd.AddCommand(setupCmd)
}
