package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wishlist",
	Short: "The SSH Directory",
	Long:  "https://github.com/charmbracelet/soft-serve",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to the config file to use")
	rootCmd.Flags().BoolP("help", "h", false, "help for wishlist")
	rootCmd.Flags().BoolP("version", "v", false, "version for wishlist")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
