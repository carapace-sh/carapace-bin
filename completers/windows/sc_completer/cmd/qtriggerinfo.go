package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qtriggerinfoCmd = &cobra.Command{
	Use:   "qtriggerinfo",
	Short: "query the trigger parameters of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qtriggerinfoCmd).Standalone()
	rootCmd.AddCommand(qtriggerinfoCmd)
}
