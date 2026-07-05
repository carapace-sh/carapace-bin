package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unsetrequestedCmd = &cobra.Command{
	Use:   "unsetrequested",
	Short: "Mark a port as unrequested",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsetrequestedCmd).Standalone()
	rootCmd.AddCommand(unsetrequestedCmd)
}
