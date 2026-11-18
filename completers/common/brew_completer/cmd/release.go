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

	releaseCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	releaseCmd.Flags().BoolP("help", "h", false, "Show this message.")
	releaseCmd.Flags().Bool("major", false, "Create a major release.")
	releaseCmd.Flags().Bool("minor", false, "Create a minor release.")
	releaseCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	releaseCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(releaseCmd)
}
