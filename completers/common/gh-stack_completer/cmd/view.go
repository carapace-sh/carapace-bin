package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the current stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCmd).Standalone()

	viewCmd.Flags().Bool("json", false, "Output stack data as JSON")
	viewCmd.Flags().BoolP("short", "s", false, "Show compact output")
	rootCmd.AddCommand(viewCmd)
}
