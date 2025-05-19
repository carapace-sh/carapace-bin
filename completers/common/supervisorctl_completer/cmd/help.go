package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help for <action>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			vals := make([]string, 0)
			for _, cmd := range helpCmd.Parent().Commands() {
				if cmd != helpCmd && !cmd.Hidden {
					vals = append(vals, cmd.Name(), cmd.Short)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}),
	)
}
