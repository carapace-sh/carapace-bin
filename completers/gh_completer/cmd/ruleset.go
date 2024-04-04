package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var rulesetCmd = &cobra.Command{
	Use:     "ruleset <command>",
	Short:   "View info about repo rulesets",
	Aliases: []string{"rs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rulesetCmd).Standalone()

	rulesetCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(rulesetCmd)

	carapace.Gen(rulesetCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
