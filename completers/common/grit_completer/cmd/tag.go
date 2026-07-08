package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "List tags, or create / delete one. A new tag points at HEAD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()

	tagCmd.Flags().BoolP("delete", "d", false, "Delete the named tag instead of creating it")
	tagCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(tagCmd)

	carapace.Gen(tagCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if tagCmd.Flag("delete").Changed {
				return git.ActionRefs(git.RefOption{Tags: true}).FilterArgs()
			}
			switch len(c.Args) {
			case 0:
				return git.ActionRefs(git.RefOption{Tags: true})
			case 1:
				return git.ActionRefs(git.RefOption{}.Default())
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
