package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-svn",
	Short: "Bidirectional operation between a Subversion repository and Git",
	Long:  "https://git-scm.com/docs/git-svn",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().StringP("id", "i", "", "Set GIT_SVN_ID")
	rootCmd.Flags().Bool("minimize-connections", false, "Minimize SVN connections")
	rootCmd.Flags().StringP("svn-remote", "R", "", "SVN remote to use")
	rootCmd.Flags().BoolP("version", "V", false, "Show version")
}
