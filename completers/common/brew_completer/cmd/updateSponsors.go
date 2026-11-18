package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateSponsorsCmd = &cobra.Command{
	Use:     "update-sponsors",
	Short:   "Update the list of GitHub Sponsors in the `Homebrew/brew` README",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateSponsorsCmd).Standalone()

	updateSponsorsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	updateSponsorsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	updateSponsorsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	updateSponsorsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateSponsorsCmd)
}
