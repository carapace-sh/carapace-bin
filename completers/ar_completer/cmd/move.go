package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:     "move",
	Aliases: []string{"m"},
	Short:   "move file(s) in the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	moveCmd.Flags().BoolS("a", "a", false, "put file(s) after [member-name")
	moveCmd.Flags().BoolS("b", "b", false, "put file(s) before [member-name")
	addGenericOptions(moveCmd)

	moveCmd.MarkFlagsMutuallyExclusive("a", "b")

	carapace.Gen(moveCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if moveCmd.Flag("a").Changed || moveCmd.Flag("b").Changed {
				switch len(c.Args) {
				case 0:
					return carapace.ActionValues()

				case 1:
					return carapace.ActionFiles()

				default:
					return fs.ActionArFileContents(c.Args[1]).Invoke(c).Filter(c.Args[1:]...).ToA()
				}
			}

			if len(c.Args) == 0 {
				return carapace.ActionFiles()
			}
			return fs.ActionArFileContents(c.Args[0]).Invoke(c).Filter(c.Args[1:]...).ToA()
		}),
	)
}
