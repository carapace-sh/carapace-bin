package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:     "pr <command>",
	Short:   "Manage pull requests",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prCmd).Standalone()
	prCmd.AddGroup(
		&cobra.Group{ID: "Targeted commands", Title: ""},
		&cobra.Group{ID: "General commands", Title: ""},
	)

	prCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(prCmd)

	carapace.Gen(prCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
