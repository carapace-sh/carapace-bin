package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateSelfCmd = &cobra.Command{
	Use:   "update-self",
	Short: "Update this kitten binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateSelfCmd).Standalone()

	updateSelfCmd.Flags().String("fetch-version", "", "The version to fetch. The special words latest and nightly fetch the latest stable and nightly release respectively. Other values can be, for example: 0.45.0.")
	updateSelfCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	rootCmd.AddCommand(updateSelfCmd)

	carapace.Gen(updateSelfCmd).FlagCompletion(carapace.ActionMap{
		"fetch-version": carapace.ActionValues("latest", "nightly"),
	})
}
