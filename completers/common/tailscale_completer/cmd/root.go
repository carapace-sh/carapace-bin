package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tailscale",
	Short: "The easiest, most secure way to use WireGuard",
	Long: "https://tailscale.com",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("socket", "", "path to tailscaled socket")
	rootCmd.Flags().Bool("json-docs", false, "print JSON-encoded docs for all subcommands and flags")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"socket": carapace.ActionFiles(),
	})
}
