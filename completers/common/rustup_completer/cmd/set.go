package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Alter rustup settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(setCmd)
}
