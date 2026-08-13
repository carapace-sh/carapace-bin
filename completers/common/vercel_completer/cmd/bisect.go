package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bisectCmd = &cobra.Command{
	Use:   "bisect",
	Short: "Bisect the current project interactively or via an automated test script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisectCmd).Standalone()

	bisectCmd.Flags().StringP("bad", "b", "", "Known bad URL")
	bisectCmd.Flags().StringP("good", "g", "", "Known good URL")
	bisectCmd.Flags().BoolP("open", "o", false, "Automatically open each URL in the browser")
	bisectCmd.Flags().StringP("path", "p", "", "Subpath of the deployment URL to test")
	bisectCmd.Flags().StringP("run", "r", "", "Test script to run for each deployment")

	rootCmd.AddCommand(bisectCmd)
}
