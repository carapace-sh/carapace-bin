package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteinprogressCmd = &cobra.Command{
	Use:   "deleteinprogress",
	Short: "delete all in-progress backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteinprogressCmd).Standalone()
	rootCmd.AddCommand(deleteinprogressCmd)
}
