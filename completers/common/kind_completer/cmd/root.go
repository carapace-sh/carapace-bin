package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kind",
	Short: "kind is a tool for managing local Kubernetes clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "silence all stderr output")
	rootCmd.PersistentFlags().StringP("verbosity", "v", "", "info log verbosity, higher value produces more output")
}
