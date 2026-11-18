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

	tapNewCmd.Flags().String("branch", "", "Initialize Git repository and setup GitHub Actions workflows with the specified branch name (default: `main`).")
	tapNewCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	tapNewCmd.Flags().Bool("github-packages", false, "Upload bottles to GitHub Packages.")
	tapNewCmd.Flags().BoolP("help", "h", false, "Show this message.")
	tapNewCmd.Flags().Bool("no-git", false, "Don't initialize a Git repository for the tap.")
	tapNewCmd.Flags().String("pull-label", "", "Label name for pull requests ready to be pulled (default: `pr-pull`).")
	tapNewCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	tapNewCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(tapNewCmd)
}
