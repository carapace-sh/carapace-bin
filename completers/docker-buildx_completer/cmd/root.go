package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-buildx",
	Short: "Docker Buildx",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().String("builder", "", "Override the configured builder instance")
	rootCmd.Flags().BoolP("help", "h", false, "help for ./buildx")

	// TODO builder completion
}
