package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "argparse",
	Short: "Parse options passed to a fish script or function",
	Long:  "https://fishshell.com/docs/current/cmds/argparse.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("exclusive", "x", "", "mutually exclusive options")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolP("ignore-unknown", "i", false, "ignore unknown options")
	rootCmd.Flags().StringP("max-args", "X", "", "maximum number of non-option arguments")
	rootCmd.Flags().StringP("min-args", "N", "", "minimum number of non-option arguments")
	rootCmd.Flags().BoolP("move-unknown", "u", false, "move unknown options to argv_opts")
	rootCmd.Flags().BoolP("name", "n", false, "name for error messages")
	rootCmd.Flags().BoolP("stop-nonopt", "s", false, "stop scanning on first non-option argument")
	rootCmd.Flags().Bool("strict-longopts", false, "require -- for long options")
	rootCmd.Flags().String("unknown-arguments", "", "parsing for unknown options")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"unknown-arguments": carapace.ActionValues("optional", "required", "none"),
	})
}
