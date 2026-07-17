package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [toolchain]",
	Short: "Set the in-use or default toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()

	useCmd.Flags().String("format", "text", "Output format (text|json)")
	useCmd.Flags().BoolP("global-default", "g", false, "Set the global default toolchain")
	useCmd.Flags().BoolP("help", "h", false, "Show help information")
	useCmd.Flags().BoolP("print-location", "p", false, "Print the location of the in-use toolchain")

	rootCmd.AddCommand(useCmd)

	carapace.Gen(useCmd).FlagCompletion(carapace.ActionMap{
		"format": actionFormat(),
	})

	carapace.Gen(useCmd).PositionalCompletion(
		actionToolchains(),
	)
}
