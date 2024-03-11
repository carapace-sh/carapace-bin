package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var work_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync workspace build list to modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_syncCmd).Standalone()
	work_syncCmd.Flags().SetInterspersed(false)

	workCmd.AddCommand(work_syncCmd)
}
