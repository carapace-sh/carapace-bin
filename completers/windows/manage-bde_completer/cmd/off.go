package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "disable BitLocker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offCmd).Standalone()
	rootCmd.AddCommand(offCmd)
}
