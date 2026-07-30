package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "display or set the time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timeCmd).Standalone()
	rootCmd.AddCommand(timeCmd)
}
