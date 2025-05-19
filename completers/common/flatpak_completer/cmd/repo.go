package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:     "repo [OPTION…] LOCATION",
	Short:   "Repository maintenance",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()

	repoCmd.Flags().Bool("branches", false, "List the branches in the repository")
	repoCmd.Flags().String("commits", "", "Show commits for a branch")
	repoCmd.Flags().BoolP("help", "h", false, "Show help options")
	repoCmd.Flags().Bool("info", false, "Print general information about the repository")
	repoCmd.Flags().String("metadata", "", "Print metadata for a branch")
	repoCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	repoCmd.Flags().Bool("subset", false, "Limit information to subsets with this prefix")
	repoCmd.Flags().Bool("subsets", false, "Print information about the repo subsets")
	repoCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(repoCmd)

	// TODO flag
	carapace.Gen(repoCmd).FlagCompletion(carapace.ActionMap{
		// "commits"
		// "metadata"
	})

	// TODO positional
}
