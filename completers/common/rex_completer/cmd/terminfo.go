package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoCmd = &cobra.Command{
	Use:   "terminfo",
	Short: "Set up and report Rex's terminal type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoCmd).Standalone()

	rootCmd.AddCommand(terminfoCmd)
}
