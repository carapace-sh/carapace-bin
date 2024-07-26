package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeQueueTableCmd = &cobra.Command{
	Use:     "make:queue-table",
	Short:   "Create a migration for the queue jobs database table",
	Aliases: []string{"queue:table"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeQueueTableCmd).Standalone()

	rootCmd.AddCommand(makeQueueTableCmd)
}
