package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var softRebootCmd = &cobra.Command{
	Use:     "soft-reboot",
	Short:   "Shut down and reboot userspace",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(softRebootCmd).Standalone()

	rootCmd.AddCommand(softRebootCmd)
}
