package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var tag_listCmd = &cobra.Command{
	Use:   "list [OPTIONS] [NAMES]...",
	Short: "List tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_listCmd).Standalone()

	tag_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tag_listCmd.Flags().StringP("template", "T", "", "Render each tag using the given template")
	tagCmd.AddCommand(tag_listCmd)

	carapace.Gen(tag_listCmd).PositionalAnyCompletion(
		jj.ActionTags(),
	)
}
