package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release <command>",
	Short:   "Manage releases",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()
	releaseCmd.AddGroup(
		&cobra.Group{ID: "General commands", Title: ""},
		&cobra.Group{ID: "Targeted commands", Title: ""},
	)

	releaseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
