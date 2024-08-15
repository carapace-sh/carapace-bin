package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showIndexCmd = &cobra.Command{
	Use:     "show-index",
	Short:   "Show packed archive index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(showIndexCmd).Standalone()

	showIndexCmd.Flags().String("object-format", "", "specify the given object format for the index file")
	rootCmd.AddCommand(showIndexCmd)

	carapace.Gen(showIndexCmd).FlagCompletion(carapace.ActionMap{
		"object-format": carapace.ActionValues("sha1", "sha256"),
	})
}
