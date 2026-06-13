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

	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("i", "i", false, "ignore unknown")
	rootCmd.Flags().String("max-args", "", "maximum number of non-option arguments")
	rootCmd.Flags().String("min-args", "", "minimum number of non-option arguments")
	rootCmd.Flags().BoolS("n", "n", false, "name for error messages")
	rootCmd.Flags().Bool("name", false, "name for error messages")
	rootCmd.Flags().BoolS("s", "s", false, "stop scanning on first non-option argument")
	rootCmd.Flags().Bool("stop-nonopt", false, "stop scanning on first non-option argument")
	rootCmd.Flags().Bool("strict-longopts", false, "require -- for long options")
	rootCmd.Flags().String("unknown-arguments", "", "parsing for unknown options")
	rootCmd.Flags().BoolS("u", "u", false, "move unknown options")
	rootCmd.Flags().Bool("move-unknown", false, "move unknown options to argv_opts")
	rootCmd.Flags().String("exclusive", "", "mutually exclusive options")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"unknown-arguments": carapace.ActionValues("optional", "required", "none"),
	})
}
