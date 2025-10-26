package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taplo",
	Short: "A TOML toolkit written in Rust",
	Long:  "https://taplo.tamasfe.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("colors", "", "")
	rootCmd.PersistentFlags().BoolS("h", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("help", false, "Print help (see a summary with '-h')")
	rootCmd.PersistentFlags().Bool("log-spans", false, "Enable logging spans")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable a verbose logging format")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"colors": carapace.ActionValuesDescribed(
			"auto", "Determine whether to colorize output automatically",
			"always", "Always colorize output",
			"never", "Never colorize output",
		),
	})
}
