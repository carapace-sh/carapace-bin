package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var tag_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete existing tags",
	Aliases: []string{"d"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_deleteCmd).Standalone()

	tag_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tagCmd.AddCommand(tag_deleteCmd)

	carapace.Gen(tag_deleteCmd).PositionalAnyCompletion(
		jj.ActionTags().FilterArgs(),
	)
}
