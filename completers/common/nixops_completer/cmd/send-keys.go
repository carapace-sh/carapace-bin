package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var SendKeysCmd = &cobra.Command{
	Use:   "send-keys",
	Short: "Send Keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(SendKeysCmd).Standalone()
	rootCmd.AddCommand(SendKeysCmd)
}
