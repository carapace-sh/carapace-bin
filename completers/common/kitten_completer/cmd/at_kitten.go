package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_kittenCmd = &cobra.Command{
	Use:   "kitten",
	Short: "Run a kitten",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_kittenCmd).Standalone()

	at_kittenCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_kittenCmd.Flags().StringP("match", "m", "", "The window to match")
	atCmd.AddCommand(at_kittenCmd)

}
