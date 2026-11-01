package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var interrogateCmd = &cobra.Command{
	Use:   "interrogate",
	Short: "send an INTERROGATE control request to a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(interrogateCmd).Standalone()
	rootCmd.AddCommand(interrogateCmd)
}
