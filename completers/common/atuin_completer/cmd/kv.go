package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Get or set small key-value pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kvCmd).Standalone()

	kvCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(kvCmd)
}
