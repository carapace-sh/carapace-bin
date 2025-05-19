package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release",
	Short:   "Create a new draft Homebrew/brew release with the appropriate version number and release notes",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()

	releaseCmd.Flags().Bool("debug", false, "Display any debugging information.")
	releaseCmd.Flags().Bool("help", false, "Show this message.")
	releaseCmd.Flags().Bool("major", false, "Create a major release.")
	releaseCmd.Flags().Bool("minor", false, "Create a minor release.")
	releaseCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	releaseCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(releaseCmd)
}
