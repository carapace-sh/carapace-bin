package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:     "cache <command>",
	Short:   "Manage Github Actions caches",
	GroupID: "actions",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	cacheCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(cacheCmd)

	carapace.Gen(cacheCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepoOverride(cacheCmd),
	})
}
