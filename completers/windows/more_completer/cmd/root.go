package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "more",
	Short: "display one screen of output at a time",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/more",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("c", "c", false, "clear screen before displaying page")
	rootCmd.Flags().BoolP("e", "e", false, "enable extended features")
	rootCmd.Flags().BoolP("p", "p", false, "expand form-feed characters")
	rootCmd.Flags().BoolP("s", "s", false, "squeeze multiple blank lines into one")
	rootCmd.Flags().BoolP("t", "t", false, "expand tabs to spaces")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
