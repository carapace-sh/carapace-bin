package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:     "reboot",
	Short:   "Shut down and reboot the system",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebootCmd).Standalone()

	rootCmd.AddCommand(rebootCmd)
}
