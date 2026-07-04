package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "choice",
	Short: "prompt the user to select one item from a list of choices",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/choice",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("n", "n", false, "do not display the choices")
	rootCmd.Flags().BoolP("s", "s", false, "case-sensitive comparison")
}
