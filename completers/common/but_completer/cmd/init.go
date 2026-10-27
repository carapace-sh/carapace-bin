package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Create an empty Git repository in the current directory",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "other commands",
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("commit", false, "Also create an initial empty commit")
	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(initCmd)
}
