package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var creditsCmd = &cobra.Command{
	Use:    "credits",
	Short:  "View credits for this tool",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(creditsCmd).Standalone()

	creditsCmd.Flags().BoolP("static", "s", false, "Print a static version of the credits")
	rootCmd.AddCommand(creditsCmd)
}
