package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount a directory from a machine via SSHFS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountCmd).Standalone()
	rootCmd.AddCommand(mountCmd)
}
