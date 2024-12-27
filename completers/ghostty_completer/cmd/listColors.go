package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listColorsCmd = &cobra.Command{
	Use:   "+list-colors",
	Short: "list all the named RGB colors in Ghostty",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listColorsCmd).Standalone()

	listColorsCmd.Flags().Bool("help", false, "show help")
	rootCmd.AddCommand(listColorsCmd)
}
