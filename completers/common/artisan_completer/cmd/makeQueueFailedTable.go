package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeQueueFailedTableCmd = &cobra.Command{
	Use:     "make:queue-failed-table",
	Short:   "Create a migration for the failed queue jobs database table",
	Aliases: []string{"queue:failed-table"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeQueueFailedTableCmd).Standalone()

	rootCmd.AddCommand(makeQueueFailedTableCmd)
}
