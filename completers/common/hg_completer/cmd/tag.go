package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Short:   "add one or more tags for the current or given revision",
	GroupID: groups[group_change_organization].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()

	tagCmd.Flags().StringP("date", "d", "", "record the specified date as commit date")
	tagCmd.Flags().BoolP("edit", "e", false, "invoke editor on commit messages")
	tagCmd.Flags().BoolP("force", "f", false, "force tag")
	tagCmd.Flags().BoolP("local", "l", false, "make the tag local")
	tagCmd.Flags().StringP("message", "m", "", "use text as commit message")
	tagCmd.Flags().Bool("remove", false, "remove a tag")
	tagCmd.Flags().StringP("rev", "r", "", "revision to tag")
	tagCmd.Flags().StringP("user", "u", "", "record the specified user as committer")
	rootCmd.AddCommand(tagCmd)

	carapace.Gen(tagCmd).FlagCompletion(carapace.ActionMap{
		"rev": hg.ActionRevisions(),
	})

	carapace.Gen(tagCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if tagCmd.Flags().Changed("remove") {
				return hg.ActionTags()
			}
			return carapace.ActionValues()
		}),
	)
}
