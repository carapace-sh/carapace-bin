package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:     "container",
	Short:   "Manage containers",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerCmd).Standalone()

	rootCmd.AddCommand(containerCmd)
}
