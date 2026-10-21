package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sendKeysCmd = &cobra.Command{
	Use:   "send-keys",
	Short: "send encryption keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendKeysCmd).Standalone()

	sendKeysCmd.Flags().StringSlice("exclude", nil, "Send keys to all except the specified machines")
	sendKeysCmd.Flags().StringSlice("include", nil, "Send keys only to the specified machines")

	rootCmd.AddCommand(sendKeysCmd)
}
