package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:     "project <command>",
	Short:   "Work with GitHub Projects.",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectCmd).Standalone()

	rootCmd.AddCommand(projectCmd)
}
