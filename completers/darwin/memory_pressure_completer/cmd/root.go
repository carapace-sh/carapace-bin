package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "memory_pressure",
	Short: "apply real or simulate memory pressure on the system",
	Long:  "https://man.freebsd.org/cgi/man.cgi?memory_pressure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "Simulate memory pressure")
	rootCmd.Flags().StringS("l", "l", "", "Pressure level (warn or critical)")
	rootCmd.Flags().StringS("p", "p", "", "Percent free")
	rootCmd.Flags().StringS("s", "s", "", "Sleep seconds")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"l": carapace.ActionValues("warn", "critical"),
	})
}
