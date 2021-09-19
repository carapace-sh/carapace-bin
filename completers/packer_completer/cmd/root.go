package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "packer",
	Short: "Create identical machine images for multiple platforms from a single source configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().Bool("v", false, "show version")
	rootCmd.Flags().Bool("version", false, "show version")

	rootCmd.PersistentFlags().Bool("h", false, "show help")
	rootCmd.PersistentFlags().Bool("help", false, "show help")
}
