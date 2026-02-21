package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display package info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("files", "f", false, "Show a list of package files")
	infoCmd.Flags().StringP("component", "c", "", "Info about the given component")
	infoCmd.Flags().Bool("files-path", false, "Show only paths")
	infoCmd.Flags().BoolP("short", "s", false, "Do not show details")
	infoCmd.Flags().Bool("xml", false, "Output in XML format")

	rootCmd.AddCommand(infoCmd)
}
