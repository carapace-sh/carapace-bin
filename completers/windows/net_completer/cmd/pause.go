package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "pause a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pauseCmd).Standalone()
	rootCmd.AddCommand(pauseCmd)
}
