package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Write arguments after expansion of pseudo-portnames and selectors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(echoCmd).Standalone()
	rootCmd.AddCommand(echoCmd)
}
