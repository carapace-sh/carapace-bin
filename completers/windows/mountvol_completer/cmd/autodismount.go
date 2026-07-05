package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autodismountCmd = &cobra.Command{
	Use:   "autodismount",
	Short: "enable automatic mounting of new volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autodismountCmd).Standalone()
	rootCmd.AddCommand(autodismountCmd)
}
