package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var tag_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Create or update tags",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_setCmd).Standalone()

	tag_setCmd.Flags().Bool("allow-move", false, "Allow moving existing tags")
	tag_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tag_setCmd.Flags().StringP("revision", "r", "", "Target revision to point to")
	tag_setCmd.Flags().String("to", "", "Target revision to point to")
	tagCmd.AddCommand(tag_setCmd)

	carapace.Gen(tag_setCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(tag_setCmd).PositionalAnyCompletion(
		jj.ActionTags().FilterArgs(),
	)
}
