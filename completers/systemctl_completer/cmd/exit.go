package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitCmd = &cobra.Command{
	Use:     "exit",
	Short:   "Request user instance or container exit",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitCmd).Standalone()

	rootCmd.AddCommand(exitCmd)
}
