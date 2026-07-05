package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processCmd = &cobra.Command{
	Use:   "process",
	Short: "process management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processCmd).Standalone()
	rootCmd.AddCommand(processCmd)
}
