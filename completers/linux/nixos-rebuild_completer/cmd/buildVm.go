package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildVmCmd = &cobra.Command{
	Use:   "build-vm",
	Short: "Build a script that starts a NixOS virtual machine with the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildVmCmd).Standalone()
	rootCmd.AddCommand(buildVmCmd)
}
