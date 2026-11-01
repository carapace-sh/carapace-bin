package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version [-j]",
	Short:   "display zfs userland utility and kernel module version",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolS("j", "j", false, "JSON output")

	rootCmd.AddCommand(versionCmd)
}
