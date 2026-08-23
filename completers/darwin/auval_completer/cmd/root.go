package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auval",
	Short: "AudioUnit validation",
	Long:  "https://keith.github.io/xcode-manpages/auval.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("32", false, "Run in 32 bit mode")
	rootCmd.Flags().BoolS("a", "a", false, "Validate all installed AudioUnits")
	rootCmd.Flags().BoolS("de", "de", false, "Detect errors")
	rootCmd.Flags().BoolS("dw", "dw", false, "Detect warnings")
	rootCmd.Flags().StringS("f", "f", "", "File to output results to")
	rootCmd.Flags().BoolS("h", "h", false, "Print help text")
	rootCmd.Flags().StringS("s", "s", "", "AudioUnit type")
	rootCmd.Flags().StringS("v", "v", "", "Validate a specific AudioUnit")
	rootCmd.Flags().BoolS("w", "w", false, "Wait for debugger to attach")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"s": carapace.ActionValues("aufx", "aumx", "aufc", "aufb", "aufd", "aumi", "aumr", "aumc", "aumx"),
	})
}
