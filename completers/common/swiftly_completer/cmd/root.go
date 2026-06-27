package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swiftly",
	Short: "A Swift toolchain installer and manager",
	Long:  "https://github.com/swiftlang/swiftly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show help information")
	rootCmd.Flags().Bool("version", false, "Show the version")
}

func actionToolchains() carapace.Action {
	return carapace.ActionValuesDescribed(
		"latest", "Most recent stable release",
		"main-snapshot", "Most recent main branch snapshot",
	)
}

func actionFormat() carapace.Action {
	return carapace.ActionValues("text", "json")
}
