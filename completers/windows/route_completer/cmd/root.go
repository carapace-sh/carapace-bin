package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "route",
	Short: "display and modify the local routing table",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("f", "f", false, "clear all gateway entries")
	rootCmd.Flags().BoolP("p", "p", false, "make route persistent across reboots")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "add a route",
			"change", "modify an existing route",
			"delete", "delete a route",
			"print", "display the routing table",
		),
	)
}
