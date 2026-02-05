package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var tag_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List tags and their targets",
	Aliases: []string{"l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_listCmd).Standalone()

	tag_listCmd.Flags().Bool("all", false, "Show all tracked and untracked remote tags including the ones whose targets are synchronized with the local tags")
	tag_listCmd.Flags().BoolP("all-remotes", "a", false, "Show all tracked and untracked remote tags including the ones whose targets are synchronized with the local tags")
	tag_listCmd.Flags().BoolP("conflicted", "c", false, "Show conflicted tags only")
	tag_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tag_listCmd.Flags().StringSliceP("revisions", "r", nil, "Show tags whose local targets are in the given revisions")
	tag_listCmd.Flags().StringSlice("sort", nil, "Sort tags based on the given key (or multiple keys)")
	tag_listCmd.Flags().StringP("template", "T", "", "Render each tag using the given template")
	tag_listCmd.Flag("all").Hidden = true
	tagCmd.AddCommand(tag_listCmd)

	carapace.Gen(tag_listCmd).FlagCompletion(carapace.ActionMap{
		"sort": carapace.ActionValues(
			"name",
			"name-",
			"author-name",
			"author-name-",
			"author-email",
			"author-email-",
			"author-date",
			"author-date-",
			"committer-name",
			"committer-name-",
			"committer-email",
			"committer-email-",
			"committer-date",
			"committer-date-",
		).UniqueList(","), // TODO filter both asc/desc
	})

	carapace.Gen(tag_listCmd).PositionalAnyCompletion(
		jj.ActionTags(),
	)
}
