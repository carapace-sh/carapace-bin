package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var work_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync workspace build list to modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_syncCmd).Standalone()

	workCmd.AddCommand(work_syncCmd)
}
