package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Output mas config and related system info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	configCmd.Flags().Bool("json", false, "Output JSON")
	rootCmd.AddCommand(configCmd)
}
