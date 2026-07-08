package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grit",
	Short: "A simple Grit-powered CLI",
	Long:  "https://grit-scm.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("filter", "", "jq-like expression applied to JSON output (requires `--json`)")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("json", false, "Emit machine-readable JSON instead of human-readable text")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionValues(),
	})
}
