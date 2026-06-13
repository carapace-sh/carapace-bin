package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "math",
	Short: "Perform mathematics calculations",
	Long:  "https://fishshell.com/docs/current/cmds/math.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("b", "b", "", "numeric base")
	rootCmd.Flags().String("base", "", "numeric base")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().StringS("m", "m", "", "scale behavior")
	rootCmd.Flags().StringS("s", "s", "", "decimal places")
	rootCmd.Flags().String("scale", "", "decimal places")
	rootCmd.Flags().String("scale-mode", "", "scale behavior")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"base":       carapace.ActionValues("hex", "octal", "16", "8"),
		"scale-mode": carapace.ActionValues("truncate", "round", "floor", "ceiling"),
	})
}
