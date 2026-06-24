package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a stash entry without reflog entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var stash_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export stash entries to a ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var stash_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import stash entries from a commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var stash_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save local modifications to a new stash (deprecated, use push)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var stash_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store a given stash entry in the reflog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_createCmd).Standalone()
	carapace.Gen(stash_exportCmd).Standalone()
	carapace.Gen(stash_importCmd).Standalone()
	carapace.Gen(stash_saveCmd).Standalone()
	carapace.Gen(stash_storeCmd).Standalone()

	stash_saveCmd.Flags().BoolP("keep-index", "k", false, "keep changed added to index")
	stash_saveCmd.Flags().StringP("message", "m", "", "set description")
	stash_storeCmd.Flags().StringP("message", "m", "", "set description")
	stash_saveCmd.Flags().BoolP("patch", "p", false, "interactively select hunks between HEAD and working tree")
	stash_exportCmd.Flags().Bool("print", false, "print the object ID instead of writing it to a ref")
	stash_saveCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	stash_storeCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	stash_saveCmd.Flags().BoolP("staged", "S", false, "only stash staged")
	stash_exportCmd.Flags().String("to-ref", "", "save the data to the given ref")

	stashCmd.AddCommand(stash_createCmd)
	stashCmd.AddCommand(stash_exportCmd)
	stashCmd.AddCommand(stash_importCmd)
	stashCmd.AddCommand(stash_saveCmd)
	stashCmd.AddCommand(stash_storeCmd)

	carapace.Gen(stash_exportCmd).FlagCompletion(carapace.ActionMap{
		"to-ref": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(stash_exportCmd).PositionalAnyCompletion(
		git.ActionStashes(),
	)

	carapace.Gen(stash_importCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(stash_storeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
