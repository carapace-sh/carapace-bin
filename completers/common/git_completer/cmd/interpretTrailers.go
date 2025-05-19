package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var interpretTrailerCmd = &cobra.Command{
	Use:     "interpret-trailers",
	Short:   "Add or parse structured information in commit messages",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(interpretTrailerCmd).Standalone()

	interpretTrailerCmd.Flags().String("if-exists", "", "action if trailer already exists")
	interpretTrailerCmd.Flags().Bool("in-place", false, "edit files in place")
	interpretTrailerCmd.Flags().Bool("no-divider", false, "do not treat --- specially")
	interpretTrailerCmd.Flags().Bool("only-input", false, "do not apply config rules")
	interpretTrailerCmd.Flags().Bool("only-trailers", false, "output only the trailers")
	interpretTrailerCmd.Flags().Bool("parse", false, "set parsing options")
	interpretTrailerCmd.Flags().String("trailer", "", "trailer(s) to add")
	interpretTrailerCmd.Flags().Bool("trim-empty", false, "trim empty trailers")
	interpretTrailerCmd.Flags().Bool("unfold", false, "join whitespace-continued values")
	interpretTrailerCmd.Flags().String("where", "", "where to place the new trailer")
	rootCmd.AddCommand(interpretTrailerCmd)

	carapace.Gen(interpretTrailerCmd).FlagCompletion(carapace.ActionMap{
		"if-exists": carapace.ActionValues("addIfDifferent", "addIfDifferentNeighbor", "add", "replace", "doNothing"),
		"trailer":   carapace.ActionValues("doNothing", "add"),
		"where":     carapace.ActionValues("after", "before", "end", "start"),
	})

	carapace.Gen(interpretTrailerCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
