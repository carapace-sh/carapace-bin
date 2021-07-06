package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glab",
	Short: "A GitLab CLI Tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.Flags().BoolP("version", "v", false, "show glab version information")

	carapace.Gen(rootCmd)
}
