package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueForgetCmd = &cobra.Command{
	Use:   "queue:forget <id>",
	Short: "Delete a failed queue job",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueForgetCmd).Standalone()

	rootCmd.AddCommand(queueForgetCmd)
}
