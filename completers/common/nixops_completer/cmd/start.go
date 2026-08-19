package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(StartCmd).Standalone()
	rootCmd.AddCommand(StartCmd)
}
