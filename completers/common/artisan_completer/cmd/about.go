package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about [--only [ONLY]] [--json]",
	Short: "Display basic information about your application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aboutCmd).Standalone()

	aboutCmd.Flags().Bool("json", false, "Output the information as JSON")
	aboutCmd.Flags().String("only", "", "The section to display")
	rootCmd.AddCommand(aboutCmd)
}
