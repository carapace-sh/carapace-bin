package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessibilityCmd = &cobra.Command{
	Use:     "accessibility",
	Short:   "Learn about GitHub CLI's accessibility experiences",
	Aliases: []string{"a11y"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessibilityCmd).Standalone()

	accessibilityCmd.Flags().BoolP("web", "w", false, "Open the GitHub Accessibility site in your browser")
	rootCmd.AddCommand(accessibilityCmd)
}
