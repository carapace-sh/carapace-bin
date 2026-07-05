package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var triggerinfoCmd = &cobra.Command{
	Use:   "triggerinfo",
	Short: "configure the trigger parameters of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(triggerinfoCmd).Standalone()
	rootCmd.AddCommand(triggerinfoCmd)
}
