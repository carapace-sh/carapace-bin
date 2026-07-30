package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sleepstudyCmd = &cobra.Command{
	Use:   "sleepstudy",
	Short: "generate a diagnostic system power transition report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sleepstudyCmd).Standalone()
	rootCmd.AddCommand(sleepstudyCmd)
}
