package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waketimersCmd = &cobra.Command{
	Use:   "waketimers",
	Short: "enumerate active wake timers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waketimersCmd).Standalone()
	rootCmd.AddCommand(waketimersCmd)
}
