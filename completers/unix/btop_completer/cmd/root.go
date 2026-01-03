package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "btop",
	Short: "A monitor of resources",
	Long:  "https://github.com/aristocratos/btop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "Path to a config file")
	rootCmd.Flags().BoolP("debug", "d", false, "Start in debug mode with additional logs and metrics")
	rootCmd.Flags().StringP("filter", "f", "", "Set an initial process filter")
	rootCmd.Flags().Bool("force-utf", false, "Override automatic UTF locale detection")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().BoolP("low-color", "l", false, "Disable true color, 256 colors only")
	rootCmd.Flags().Bool("no-tty", false, "Force disable tty mode")
	rootCmd.Flags().StringP("preset", "p", "", "Start with a preset (0-9)")
	rootCmd.Flags().BoolP("tty", "t", false, "Force tty mode with ANSI graph symbols and 16 colors only")
	rootCmd.Flags().StringP("update", "u", "", "Set an initial update rate in milliseconds")
	rootCmd.Flags().BoolP("version", "V", false, "Show a version message and exit (more with --version)")

	// TODO filter
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"preset": number.ActionRange(number.RangeOpts{Start: 0, End: 9}),
	})
}
