package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"lookup"},
	Short:   "Output app info from the App Store",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	infoCmd.Flags().Bool("json", false, "Output JSON")
	rootCmd.AddCommand(infoCmd)
}
