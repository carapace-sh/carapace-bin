package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Build a version of golangci-lint with custom linters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customCmd).Standalone()

	customCmd.PersistentFlags().String("destination", "", "The directory path used to store the custom binary")
	customCmd.PersistentFlags().String("name", "", "The name of the custom binary")
	customCmd.PersistentFlags().String("version", "", "The golangci-lint version used to build the custom binary")
	rootCmd.AddCommand(customCmd)

	carapace.Gen(customCmd).FlagCompletion(carapace.ActionMap{
		"destination": carapace.ActionDirectories(),
	})
}
