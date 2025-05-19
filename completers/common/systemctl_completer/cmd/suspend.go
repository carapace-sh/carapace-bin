package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:     "suspend",
	Short:   "Suspend the system",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendCmd).Standalone()

	rootCmd.AddCommand(suspendCmd)
}
