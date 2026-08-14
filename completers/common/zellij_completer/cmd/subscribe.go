package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to pane render updates (viewport and scrollback)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subscribeCmd).Standalone()

	subscribeCmd.Flags().Bool("ansi", false, "Preserve ANSI styling in the output")
	subscribeCmd.Flags().StringP("format", "f", "raw", "Output format")
	subscribeCmd.Flags().BoolP("help", "h", false, "Print help")
	subscribeCmd.Flags().StringSliceP("pane-id", "p", nil, "Pane ID(s) to subscribe to (e.g. terminal_1, plugin_2, or bare number like 1)")
	subscribeCmd.Flags().StringP("scrollback", "s", "", "Include scrollback lines in initial delivery. Bare --scrollback = all scrollback, --scrollback N = last N lines")
	subscribeCmd.MarkFlagRequired("pane-id")
	rootCmd.AddCommand(subscribeCmd)

	carapace.Gen(subscribeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("raw", "json"),
	})
}
