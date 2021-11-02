package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doctl",
	Short: "doctl is a command line interface (CLI) for the DigitalOcean API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().StringP("access-token", "t", "", "API V2 access token")
	rootCmd.PersistentFlags().StringP("api-url", "u", "", "Override default API endpoint")
	rootCmd.PersistentFlags().StringP("config", "c", "/home/rsteube/.config/doctl/config.yaml", "Specify a custom config file")
	rootCmd.PersistentFlags().String("context", "", "Specify a custom authentication context name")
	rootCmd.Flags().BoolP("help", "h", false, "help for doctl")
	rootCmd.PersistentFlags().StringP("output", "o", "text", "Desired output format [text|json]")
	rootCmd.PersistentFlags().Bool("trace", false, "Show a log of network activity while performing a command")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"output": carapace.ActionValues("text", "json"),
	})
}
