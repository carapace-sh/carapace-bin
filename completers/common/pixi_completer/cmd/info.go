package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information about the system, workspace and environments for the current machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("extended", false, "Show cache and environment size")
	infoCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	rootCmd.AddCommand(infoCmd)
}
