package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:     "print",
	Aliases: []string{"p"},
	Short:   "print file(s) found in the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printCmd).Standalone()

	addGenericOptions(printCmd)

	carapace.Gen(printCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(printCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return fs.ActionArFileContents(c.Args[0]).Filter(c.Args[1:]...)
		}),
	)
}
