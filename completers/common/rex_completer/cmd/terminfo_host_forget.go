package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoHostForgetCmd = &cobra.Command{
	Use:   "forget HOST",
	Short: "Forget that a host has Rex's terminfo entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoHostForgetCmd).Standalone()

	terminfoHostCmd.AddCommand(terminfoHostForgetCmd)
}
