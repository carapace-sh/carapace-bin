package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var distcheckCmd = &cobra.Command{
	Use:   "distcheck",
	Short: "Check if distfiles have not changed and can be fetched",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distcheckCmd).Standalone()
	rootCmd.AddCommand(distcheckCmd)
}
