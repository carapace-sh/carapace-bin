package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the staged state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_commitCmd).Standalone()

	revision_commitCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_commitCmd.Flags().String("layer", "", "Commit only changes in this layer (mount path relative to repo root)")
	revision_commitCmd.Flags().StringSlice("layer-message", nil, "Per-layer commit message. Takes two values: <path> <message>. Can be specified multiple times")
	revision_commitCmd.Flags().String("link", "", "Commit only changes in this linked repository (mount path relative to repo root)")
	revision_commitCmd.Flags().StringSlice("link-message", nil, "Per-link commit message. Takes two values: <path> <message>. Can be specified multiple times")
	revision_commitCmd.Flags().Bool("stats", false, "Print stats")
	revisionCmd.AddCommand(revision_commitCmd)

	carapace.Gen(revision_commitCmd).FlagCompletion(carapace.ActionMap{
		"layer":         carapace.ActionValues(),
		"layer-message": carapace.ActionFiles(),
		"link":          carapace.ActionValues(),
		"link-message":  carapace.ActionFiles(),
	})

	carapace.Gen(revision_commitCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
