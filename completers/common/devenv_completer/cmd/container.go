package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Build, copy, or run a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerCmd).Standalone()

	rootCmd.AddCommand(containerCmd)
}
