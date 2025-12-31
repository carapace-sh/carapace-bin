package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var booCmd = &cobra.Command{
	Use:   "+boo",
	Short: "The `boo` command is used to display the animation from the Ghostty website in the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(booCmd).Standalone()

	rootCmd.AddCommand(booCmd)
}
