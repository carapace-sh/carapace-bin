package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fury",
	Short: "Command line interface to Gemfury API",
	Long:  "https://github.com/gemfury/cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("account", "a", "", "Current account username")
	rootCmd.PersistentFlags().String("api-token", "", "Inline authentication token")
	rootCmd.Flags().BoolP("help", "h", false, "help for fury")
	rootCmd.Flags().BoolP("version", "v", false, "version for fury")
}
