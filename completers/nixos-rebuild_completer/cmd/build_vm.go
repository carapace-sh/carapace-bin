package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_vmCmd = &cobra.Command{
	Use:   "build-vm",
	Short: "Build a script that starts a NixOS virtual machine with the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_vmCmd).Standalone()
	rootCmd.AddCommand(build_vmCmd)
}
