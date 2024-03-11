package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mktagCmd = &cobra.Command{
	Use:     "mktag",
	Short:   "Creates a tag object",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(mktagCmd).Standalone()

	mktagCmd.Flags().Bool("strict", false, "enable more strict checking")
	rootCmd.AddCommand(mktagCmd)
}
