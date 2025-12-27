package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timespanCmd = &cobra.Command{
	Use:   "timespan",
	Short: "Validate a time span",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timespanCmd).Standalone()

	rootCmd.AddCommand(timespanCmd)
}
