package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:     "replace",
	Aliases: []string{"r"},
	Short:   "replace existing or insert new file(s) into the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replaceCmd).Standalone()

	replaceCmd.Flags().BoolS("a", "a", false, "put file(s) after [member-name")
	replaceCmd.Flags().BoolS("b", "b", false, "put file(s) before [member-name")
	replaceCmd.Flags().BoolS("f", "f", false, "truncate inserted file names")
	replaceCmd.Flags().BoolS("u", "u", false, "nly replace files that are newer than current archive contents")
	addGenericOptions(replaceCmd)

	carapace.Gen(replaceCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
