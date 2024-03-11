package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:     "bump",
	Short:   "Display out-of-date brew formulae and the latest version available",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpCmd).Standalone()

	bumpCmd.Flags().Bool("cask", false, "Check only casks.")
	bumpCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bumpCmd.Flags().Bool("formula", false, "Check only formulae.")
	bumpCmd.Flags().Bool("full-name", false, "Print formulae/casks with fully-qualified names.")
	bumpCmd.Flags().Bool("help", false, "Show this message.")
	bumpCmd.Flags().Bool("installed", false, "Check formulae and casks that are currently installed.")
	bumpCmd.Flags().Bool("limit", false, "Limit number of package results returned.")
	bumpCmd.Flags().Bool("no-pull-requests", false, "Do not retrieve pull requests from GitHub.")
	bumpCmd.Flags().Bool("open-pr", false, "Open a pull request for the new version if none have been opened yet.")
	bumpCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bumpCmd.Flags().Bool("start-with", false, "Letter or word that the list of package results should alphabetically follow.")
	bumpCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(bumpCmd)
}
