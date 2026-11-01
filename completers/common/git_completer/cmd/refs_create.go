package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var refs_createCmd = &cobra.Command{
	Use:   "create <ref> <new-value>",
	Short: "Create a reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_createCmd).Standalone()

	refs_createCmd.Flags().Bool("create-reflog", false, "create a reflog")
	refs_createCmd.Flags().String("message", "", "reason of the update")
	refs_createCmd.Flags().Bool("no-deref", false, "update <refname> not the one it points to")
	refsCmd.AddCommand(refs_createCmd)

	carapace.Gen(refs_createCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
