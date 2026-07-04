package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listKeychainsCmd = &cobra.Command{
	Use:   "list-keychains",
	Short: "Display or manipulate the keychain search list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeychainsCmd).Standalone()
	listKeychainsCmd.Flags().StringP("domain", "d", "", "Use the specified preference domain")
	listKeychainsCmd.Flags().StringP("set", "s", "", "Set the search list to the specified keychains")
	rootCmd.AddCommand(listKeychainsCmd)
}
