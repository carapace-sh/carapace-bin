package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmesg",
	Short: "display the system message buffer",
	Long:  "https://keith.github.io/xcode-manpages/dmesg.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("core", "M", "", "Extract values associated with the name list from the specified core")
	rootCmd.Flags().StringP("system", "N", "", "Extract the name list from the specified system")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"core":   carapace.ActionFiles(),
		"system": carapace.ActionFiles(),
	})
}
