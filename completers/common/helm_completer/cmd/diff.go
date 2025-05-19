package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:     "diff",
	Short:   "Preview helm upgrade changes as a diff",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	rootCmd.AddCommand(diffCmd)
}
