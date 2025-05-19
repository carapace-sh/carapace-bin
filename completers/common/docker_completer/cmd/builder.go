package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var builderCmd = &cobra.Command{
	Use:     "builder",
	Short:   "Manage builds",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(builderCmd).Standalone()
	rootCmd.AddCommand(builderCmd)
}
