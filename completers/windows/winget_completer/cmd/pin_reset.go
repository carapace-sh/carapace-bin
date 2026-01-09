package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pin_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset pins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pin_resetCmd).Standalone()

	pin_resetCmd.Flags().Bool("force", false, "Direct run the command and continue with non security related issues")
	pin_resetCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	pinCmd.AddCommand(pin_resetCmd)

	// TODO flag completion
}
