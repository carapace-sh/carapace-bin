package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export stash entries to a ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_exportCmd).Standalone()

	stash_exportCmd.Flags().Bool("print", false, "print the object ID instead of writing it to a ref")
	stash_exportCmd.Flags().String("to-ref", "", "save the data to the given ref")
	stashCmd.AddCommand(stash_exportCmd)

	carapace.Gen(stash_exportCmd).FlagCompletion(carapace.ActionMap{
		"to-ref": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(stash_exportCmd).PositionalAnyCompletion(
		git.ActionStashes(),
	)
}
