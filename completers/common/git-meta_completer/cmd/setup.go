package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initialize a metadata remote from a project-local `.git-meta` file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(setupCmd)
}
