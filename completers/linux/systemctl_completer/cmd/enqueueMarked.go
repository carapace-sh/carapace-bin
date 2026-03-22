package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enqueueMarkedCmd = &cobra.Command{
	Use:     "enqueue-marked",
	Short:   "Enqueue jobs for all marked units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enqueueMarkedCmd).Standalone()

	rootCmd.AddCommand(enqueueMarkedCmd)
}
