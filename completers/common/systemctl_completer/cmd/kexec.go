package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kexecCmd = &cobra.Command{
	Use:     "kexec",
	Short:   "Shut down and reboot the system with kexec",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kexecCmd).Standalone()

	rootCmd.AddCommand(kexecCmd)
}
