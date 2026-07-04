package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setrequestedCmd = &cobra.Command{
	Use:   "setrequested",
	Short: "Mark a port as requested",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setrequestedCmd).Standalone()
	rootCmd.AddCommand(setrequestedCmd)
}
