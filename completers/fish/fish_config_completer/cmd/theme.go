package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "View and choose color themes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeCmd).Standalone()

	rootCmd.AddCommand(themeCmd)
}
