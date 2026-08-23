package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteinprogressCmd)
}

var deleteinprogressCmd = &cobra.Command{
	Use:   "deleteinprogress",
	Short: "delete all in-progress backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteinprogressCmd).Standalone()
}