package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify checksums for a single repo revision from cache or a local directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_verifyCmd).Standalone()

	cache_verifyCmd.Flags().String("cache-dir", "", "Cache directory to use when verifying files from cache")
	cache_verifyCmd.Flags().Bool("fail-on-extra-files", false, "Fail if some files exist locally but are not present on the remote revision")
	cache_verifyCmd.Flags().Bool("fail-on-missing-files", false, "Fail if some files exist on the remote but are missing locally")
	cache_verifyCmd.Flags().String("format", "", "Output format")
	cache_verifyCmd.Flags().BoolP("help", "h", false, "Show help")
	cache_verifyCmd.Flags().Bool("json", false, "JSON output")
	cache_verifyCmd.Flags().String("local-dir", "", "Verify files under this directory instead of the cache")
	cache_verifyCmd.Flags().Bool("no-truncate", false, "Do not truncate values in human tables")
	cache_verifyCmd.Flags().BoolP("quiet", "q", false, "Quiet output")
	cache_verifyCmd.Flags().String("repo-type", "", "The type of repository")
	cache_verifyCmd.Flags().String("revision", "", "Git revision id (branch, tag or commit hash)")
	cache_verifyCmd.Flags().String("token", "", "A User Access Token")
	cache_verifyCmd.Flags().String("type", "", "The type of repository")
	cacheCmd.AddCommand(cache_verifyCmd)

	repoTypes := carapace.ActionValues("dataset", "kernel", "model", "space")
	carapace.Gen(cache_verifyCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"format":    carapace.ActionValues("agent", "auto", "human", "json", "quiet"),
		"local-dir": carapace.ActionDirectories(),
		"repo-type": repoTypes,
		"type":      repoTypes,
	})

	carapace.Gen(cache_verifyCmd).PositionalCompletion(
		action.ActionCachedRepoNames(cache_verifyCmd),
	)
}
