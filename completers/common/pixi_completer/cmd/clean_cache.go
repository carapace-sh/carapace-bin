package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clean_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Clean the cache of your system which are touched by pixi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clean_cacheCmd).Standalone()

	clean_cacheCmd.Flags().Bool("assume-yes", false, "Answer yes to all questions")
	clean_cacheCmd.Flags().Bool("build", false, "Clean only the build related cache")
	clean_cacheCmd.Flags().Bool("build-backends", false, "Clean only the build backends environments cache")
	clean_cacheCmd.Flags().Bool("conda", false, "Clean only the conda related cache")
	clean_cacheCmd.Flags().Bool("exec", false, "Clean only `exec` cache")
	clean_cacheCmd.Flags().Bool("mapping", false, "Clean only the mapping cache")
	clean_cacheCmd.Flags().Bool("pypi", false, "Clean only the pypi related cache")
	clean_cacheCmd.Flags().Bool("repodata", false, "Clean only the repodata cache")
	clean_cacheCmd.Flags().BoolP("yes", "y", false, "Answer yes to all questions")
	clean_cacheCmd.Flag("assume-yes").Hidden = true
	cleanCmd.AddCommand(clean_cacheCmd)
}
