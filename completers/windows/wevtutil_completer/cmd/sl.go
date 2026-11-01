package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var slCmd = &cobra.Command{
	Use:   "sl",
	Short: "set event log configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(slCmd).Standalone()
	rootCmd.AddCommand(slCmd)
}
