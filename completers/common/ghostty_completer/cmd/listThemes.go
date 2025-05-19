package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listThemesCmd = &cobra.Command{
	Use:   "+list-themes",
	Short: "preview or list all the available themes for Ghostty",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listThemesCmd).Standalone()

	listThemesCmd.Flags().Bool("help", false, "show help")
	listThemesCmd.Flags().Bool("path", false, "Show the full path to the theme")
	listThemesCmd.Flags().Bool("plain", false, "Force a plain listing of themes")
	rootCmd.AddCommand(listThemesCmd)
}
