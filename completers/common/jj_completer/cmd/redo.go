package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redoCmd = &cobra.Command{
	Use:   "redo",
	Short: "Redo the most recently undone operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redoCmd).Standalone()

	redoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(redoCmd)
}
