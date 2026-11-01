package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var noserveCmd = &cobra.Command{
	Use:   "noserve",
	Short: "disable serving of mount points",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(noserveCmd).Standalone()
	rootCmd.AddCommand(noserveCmd)
}
