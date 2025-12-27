package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:     "cancel",
	Short:   "Cancel all, one, or more jobs",
	GroupID: "job",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cancelCmd).Standalone()

	rootCmd.AddCommand(cancelCmd)

	// TODO job completion
}
