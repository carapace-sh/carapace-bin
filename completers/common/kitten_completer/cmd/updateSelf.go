package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateSelfCmd = &cobra.Command{
	Use:   "update-self",
	Short: "Update this kitten binary",
	Long:  "Update this kitten binary in place to the latest available version.",
}

func init() {
	rootCmd.AddCommand(updateSelfCmd)
	carapace.Gen(updateSelfCmd).Standalone()

	updateSelfCmd.Flags().String("fetch-version", "", "The version to fetch. The special words latest and nightly fetch the latest stable and nightly release respectively. Other values can be, for example: 0.45.0.")
	updateSelfCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(updateSelfCmd).FlagCompletion(carapace.ActionMap{
		"fetch-version": carapace.ActionValues("latest", "nightly"),
	})
}
