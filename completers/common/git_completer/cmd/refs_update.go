package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var refs_updateCmd = &cobra.Command{
	Use:   "update <ref> <new-value> [<old-value>]",
	Short: "Update a reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_updateCmd).Standalone()

	refs_updateCmd.Flags().Bool("create-reflog", false, "create a reflog")
	refs_updateCmd.Flags().String("message", "", "reason of the update")
	refs_updateCmd.Flags().Bool("no-deref", false, "update <refname> not the one it points to")
	refsCmd.AddCommand(refs_updateCmd)

	carapace.Gen(refs_updateCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefObjectIDs(c.Args[0])
		}),
	)
}
