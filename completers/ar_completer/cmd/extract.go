package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:     "extract",
	Aliases: []string{"x"},
	Short:   "extract file(s) from the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractCmd).Standalone()

	extractCmd.Flags().BoolS("o", "o", false, "preserve original dates")
	addGenericOptions(extractCmd)

	carapace.Gen(extractCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(extractCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return fs.ActionArFileContents(c.Args[0]).Invoke(c).Filter(c.Args[1:]...).ToA()
		}),
	)
}
