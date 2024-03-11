package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rev_parseCmd = &cobra.Command{
	Use:     "rev-parse",
	Short:   "Pick out and massage parameters",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(rev_parseCmd).Standalone()

	rootCmd.AddCommand(rev_parseCmd)
}
