package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_gotoLayoutCmd = &cobra.Command{
	Use:   "goto-layout",
	Short: "Set the window layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_gotoLayoutCmd).Standalone()

	at_gotoLayoutCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_gotoLayoutCmd.Flags().StringP("match", "m", "", "The tab to match")
	atCmd.AddCommand(at_gotoLayoutCmd)
}
