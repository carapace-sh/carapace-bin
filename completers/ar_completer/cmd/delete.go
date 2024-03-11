package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "delete file(s) from the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	addGenericOptions(deleteCmd)

	carapace.Gen(deleteCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return fs.ActionArFileContents(c.Args[0]).Invoke(c).Filter(c.Args[1:]...).ToA()
		}),
	)
}
