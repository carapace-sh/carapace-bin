package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consumer_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete consumer(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consumer_deleteCmd).Standalone()

	consumerCmd.AddCommand(consumer_deleteCmd)
}
