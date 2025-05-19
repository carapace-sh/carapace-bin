package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goreleaser",
	Short: "Deliver Go binaries as fast and easily as possible",
	Long:  "https://goreleaser.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("debug", false, "Enable verbose mode")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose mode")
	rootCmd.Flags().BoolP("version", "v", false, "version for goreleaser")
}
