package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queryTerminalCmd = &cobra.Command{
	Use:   "query-terminal",
	Short: "Query the terminal for various capabilities",
}

func init() {
	rootCmd.AddCommand(queryTerminalCmd)
	carapace.Gen(queryTerminalCmd).Standalone()

	queryTerminalCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	queryTerminalCmd.Flags().Float64("wait-for", 0, "The amount of time (in seconds) to wait for a response from the terminal, after querying it.")

	carapace.Gen(queryTerminalCmd).PositionalAnyCompletion(carapace.ActionValuesDescribed(
		"name", "Terminal name (e.g. xterm-kitty)",
		"version", "Terminal version (e.g. 0.45.0)",
		"allow_hyperlinks", "The config option allow_hyperlinks in kitty.conf for allowing hyperlinks can be yes, no or ask",
		"font_family", "The current font's PostScript name",
		"bold_font", "The current bold font's PostScript name",
		"italic_font", "The current italic font's PostScript name",
		"bold_italic_font", "The current bold-italic font's PostScript name",
		"font_size", "The current font size in pts",
		"dpi_x", "The current DPI on the x-axis",
		"dpi_y", "The current DPI on the y-axis",
		"foreground", "The current foreground color as a 24-bit # color code",
		"background", "The current background color as a 24-bit # color code",
		"background_opacity", "The current background opacity as a number between 0 and 1",
		"clipboard_control", "The config option clipboard_control in kitty.conf for allowing reads/writes to/from the clipboard",
		"os_name", "The name of the OS the terminal is running on. kitty returns values: bsd, linux, macos, unknown",
	))
}
