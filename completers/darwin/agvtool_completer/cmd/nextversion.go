package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nextVersionCmd = &cobra.Command{
	Use:   "next-version",
	Short: "Increment version number to the next highest integral value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextVersionCmd).Standalone()
	nextVersionCmd.Flags().Bool("all", false, "Also update CFBundleVersion in Info.plist")
	rootCmd.AddCommand(nextVersionCmd)
}
