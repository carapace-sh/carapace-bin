package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:     "tags",
	Short:   "list repository tags",
	GroupID: groups[group_change_organization].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagsCmd).Standalone()

	tagsCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(tagsCmd)
}
