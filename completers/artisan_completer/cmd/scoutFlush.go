package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scoutFlushCmd = &cobra.Command{
	Use:   "scout:flush <model>",
	Short: "Flush all of the model's records from the index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scoutFlushCmd).Standalone()

	rootCmd.AddCommand(scoutFlushCmd)
}
