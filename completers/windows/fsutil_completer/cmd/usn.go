package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usnCmd = &cobra.Command{
	Use:   "usn",
	Short: "USN journal management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usnCmd).Standalone()
	rootCmd.AddCommand(usnCmd)
}
