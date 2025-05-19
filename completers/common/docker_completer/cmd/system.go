package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemCmd = &cobra.Command{
	Use:     "system",
	Short:   "Manage Docker",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemCmd).Standalone()

	rootCmd.AddCommand(systemCmd)
}
