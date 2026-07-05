package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "enable BitLocker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(onCmd).Standalone()
	rootCmd.AddCommand(onCmd)
}
