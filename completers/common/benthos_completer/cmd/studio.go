package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var studioCmd = &cobra.Command{
	Use:   "studio",
	Short: "Interact with Benthos studio (https://studio.benthos.dev)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(studioCmd).Standalone()

	studioCmd.Flags().StringP("endpoint", "e", "", "Specify the URL of the Benthos studio server to connect to.")
	rootCmd.AddCommand(studioCmd)
}
