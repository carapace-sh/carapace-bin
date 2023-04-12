package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mailinfoCmd = &cobra.Command{
	Use:     "mailinfo",
	Short:   "Extracts patch and authorship from a single e-mail message",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(mailinfoCmd).Standalone()

	rootCmd.AddCommand(mailinfoCmd)
}
