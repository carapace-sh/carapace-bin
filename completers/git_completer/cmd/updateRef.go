package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var updateRefCmd = &cobra.Command{
	Use:     "update-ref",
	Short:   "Update the object name stored in a ref safely",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(updateRefCmd).Standalone()

	updateRefCmd.Flags().Bool("create-reflog", false, "create a reflog")
	updateRefCmd.Flags().BoolS("d", "d", false, "delete the reference")
	updateRefCmd.Flags().Bool("deref", false, "opposite of --no-deref")
	updateRefCmd.Flags().StringS("m", "m", "", "reason of the update")
	updateRefCmd.Flags().Bool("no-create-reflog", false, "do not create a reflog")
	updateRefCmd.Flags().Bool("no-deref", false, "update <refname> not the one it points to")
	updateRefCmd.Flags().Bool("no-stdin", false, "do not read updates from stdin")
	updateRefCmd.Flags().Bool("stdin", false, "read updates from stdin")
	updateRefCmd.Flags().BoolS("z", "z", false, "stdin has NUL-terminated arguments")
	rootCmd.AddCommand(updateRefCmd)

	carapace.Gen(updateRefCmd).PositionalCompletion(

		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if updateRefCmd.Flag("stdin").Changed {
				return carapace.ActionValues()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)
}
