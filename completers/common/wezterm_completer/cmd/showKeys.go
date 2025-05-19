package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showKeysCmd = &cobra.Command{
	Use:   "show-keys [OPTIONS]",
	Short: "Show key assignments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showKeysCmd).Standalone()

	showKeysCmd.Flags().BoolP("help", "h", false, "Print help")
	showKeysCmd.Flags().String("key-table", "", "In lua mode, show only the named key table")
	showKeysCmd.Flags().Bool("lua", false, "Show the keys as lua config statements")
	rootCmd.AddCommand(showKeysCmd)
}
