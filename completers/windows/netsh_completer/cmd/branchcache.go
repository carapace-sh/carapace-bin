package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branchcacheCmd = &cobra.Command{
	Use:   "branchcache",
	Short: "BranchCache configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchcacheCmd).Standalone()
	rootCmd.AddCommand(branchcacheCmd)
}
