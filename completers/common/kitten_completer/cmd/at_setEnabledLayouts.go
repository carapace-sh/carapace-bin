package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setEnabledLayoutsCmd = &cobra.Command{
	Use:   "set-enabled-layouts",
	Short: "Set the enabled layouts in tabs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setEnabledLayoutsCmd).Standalone()

	at_setEnabledLayoutsCmd.Flags().Bool("configured", false, "Change the default enabled layout value so that the new value takes effect for all newly created tabs as well")
	at_setEnabledLayoutsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setEnabledLayoutsCmd.Flags().StringP("match", "m", "", "The tab to match")
	atCmd.AddCommand(at_setEnabledLayoutsCmd)

}
