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

	updateSponsorsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updateSponsorsCmd.Flags().Bool("help", false, "Show this message.")
	updateSponsorsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateSponsorsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateSponsorsCmd)
}
