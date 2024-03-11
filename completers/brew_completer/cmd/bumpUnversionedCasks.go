package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpUnversionedCasksCmd = &cobra.Command{
	Use:     "bump-unversioned-casks",
	Short:   "Check all casks with unversioned URLs in a given <tap> for updates",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpUnversionedCasksCmd).Standalone()

	bumpUnversionedCasksCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bumpUnversionedCasksCmd.Flags().Bool("dry-run", false, "Do everything except caching state and opening pull requests.")
	bumpUnversionedCasksCmd.Flags().Bool("help", false, "Show this message.")
	bumpUnversionedCasksCmd.Flags().Bool("limit", false, "Maximum runtime in minutes.")
	bumpUnversionedCasksCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bumpUnversionedCasksCmd.Flags().Bool("state-file", false, "File for caching state.")
	bumpUnversionedCasksCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(bumpUnversionedCasksCmd)
}
