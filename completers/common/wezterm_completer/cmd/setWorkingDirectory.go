package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setWorkingDirectoryCmd = &cobra.Command{
	Use:   "set-working-directory [CWD] [HOST]",
	Short: "Advise the terminal of the current working directory by emitting an OSC 7 escape sequence",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setWorkingDirectoryCmd).Standalone()

	setWorkingDirectoryCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(setWorkingDirectoryCmd)

	carapace.Gen(setWorkingDirectoryCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
