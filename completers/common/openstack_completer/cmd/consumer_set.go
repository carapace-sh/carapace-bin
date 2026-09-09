package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consumer_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set consumer properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consumer_setCmd).Standalone()

	consumer_setCmd.Flags().String("description", "", "New consumer description")
	consumerCmd.AddCommand(consumer_setCmd)
}
