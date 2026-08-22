package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "display information about an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().BoolP("verbose", "v", false, "show full paths and URLs")
	rootCmd.AddCommand(infoCmd)
}
