package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:     "auth",
	Short:   "Inspect authorization",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authCmd).Standalone()

	rootCmd.AddCommand(authCmd)
}
