package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var refs_deleteCmd = &cobra.Command{
	Use:   "delete <ref> [<old-value>]",
	Short: "Delete a reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_deleteCmd).Standalone()

	refs_deleteCmd.Flags().String("message", "", "reason of the update")
	refs_deleteCmd.Flags().Bool("no-deref", false, "update <refname> not the one it points to")
	refsCmd.AddCommand(refs_deleteCmd)

	carapace.Gen(refs_deleteCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefObjectIDs(c.Args[0])
		}),
	)
}
