package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tapNewCmd = &cobra.Command{
	Use:     "tap-new",
	Short:   "Generate the template files for a new tap",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tapNewCmd).Standalone()

	tapNewCmd.Flags().Bool("branch", false, "Initialize Git repository and setup GitHub Actions workflows with the specified branch name (default: `main`).")
	tapNewCmd.Flags().Bool("debug", false, "Display any debugging information.")
	tapNewCmd.Flags().Bool("github-packages", false, "Upload bottles to GitHub Packages.")
	tapNewCmd.Flags().Bool("help", false, "Show this message.")
	tapNewCmd.Flags().Bool("no-git", false, "Don't initialize a Git repository for the tap.")
	tapNewCmd.Flags().Bool("pull-label", false, "Label name for pull requests ready to be pulled (default: `pr-pull`).")
	tapNewCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	tapNewCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(tapNewCmd)
}
