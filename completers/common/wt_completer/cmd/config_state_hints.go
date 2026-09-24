package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_hintsCmd = &cobra.Command{
	Use:    "hints",
	Short:  "One-time hints shown in this repo",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_hintsCmd).Standalone()

	config_state_hintsCmd.PersistentFlags().String("format", "text", "Output format (text, json) [default: text]")
	config_state_hintsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_hintsCmd)

	carapace.Gen(config_state_hintsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
	})
}
