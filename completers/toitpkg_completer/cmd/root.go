package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toitpkg",
	Short: "Run pkg commands",
	Long:  "https://github.com/toitlang/tpkg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().String("cache", "", "cache dir")
	rootCmd.PersistentFlags().String("config", "", "config file")
	rootCmd.Flags().BoolP("help", "h", false, "help for toitpkg")
	rootCmd.PersistentFlags().Bool("no-autosync", false, "Don't automatically sync")
	rootCmd.PersistentFlags().Bool("no-default-registry", false, "Don't use default registry if none exists")
	rootCmd.PersistentFlags().String("sdk-version", "", "The SDK version")
	rootCmd.PersistentFlags().Bool("track", false, "Print tracking information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache":  carapace.ActionDirectories(),
		"config": carapace.ActionFiles(".yaml"),
	})
}
