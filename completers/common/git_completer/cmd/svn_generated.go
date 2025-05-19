package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var svnCmd = &cobra.Command{
	Use:     "svn",
	Short:   "Bidirectional operation between a Subversion repository and Git",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(svnCmd).Standalone()

	rootCmd.AddCommand(svnCmd)
}
