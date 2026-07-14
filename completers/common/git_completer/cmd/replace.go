package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:     "replace",
	Short:   "Create, list, delete refs to replace objects",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(replaceCmd).Standalone()

	replaceCmd.Flags().Bool("convert-graft-file", false, "creates graft commits for all entries in $GIT_DIR/info/grafts and deletes that file upon success")
	replaceCmd.Flags().BoolP("delete", "d", false, "delete existing replace refs for the given objects")
	replaceCmd.Flags().String("edit", "", "edit an object’s content interactively")
	replaceCmd.Flags().BoolP("force", "f", false, "if an existing replace ref for the same object exists, it will be overwritten")
	replaceCmd.Flags().String("format", "", "when listing, use the specified <format>")
	replaceCmd.Flags().String("graft", "", "create a graft commit")
	replaceCmd.Flags().StringP("list", "l", "", "list replace refs for objects that match the given pattern")
	replaceCmd.Flags().Bool("raw", false, "when editing, provide the raw object contents rather than pretty-printed ones")
	rootCmd.AddCommand(replaceCmd)

	replaceCmd.MarkFlagsMutuallyExclusive(
		"convert-graft-file",
		"delete",
		"edit",
		"graft",
	)

	carapace.Gen(replaceCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"short", "<replaced sha1>",
			"medium", "<replaced sha1> → <replacement sha1>",
			"long", "<replaced sha1> (<replaced type>) → <replacement sha1> (<replacement type>)",
		),
	})

	// TODO positional completion
}
