package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gsa <file> [<diff file>]",
	Short: "A tool for analyzing the size of compiled Go binaries",
	Long:  "https://github.com/Zxilly/go-size-analyzer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("compact", false, "Hide function details, replacement with size")
	rootCmd.Flags().StringP("format", "f", "", "Output format")
	rootCmd.Flags().String("height", "", "Height of the svg treemap")
	rootCmd.Flags().BoolP("help", "h", false, "Show context-sensitive help")
	rootCmd.Flags().Bool("hide-main", false, "Hide main package")
	rootCmd.Flags().Bool("hide-sections", false, "Hide sections")
	rootCmd.Flags().Bool("hide-std", false, "Hide standard library")
	rootCmd.Flags().Bool("imports", false, "Try analyze package imports from source")
	rootCmd.Flags().String("indent", "", "Indentation for json output")
	rootCmd.Flags().String("listen", "", "listen address")
	rootCmd.Flags().String("margin-box", "", "Margin between boxes")
	rootCmd.Flags().Bool("no-disasm", false, "Skip disassembly pass")
	rootCmd.Flags().Bool("no-dwarf", false, "Skip dwarf pass")
	rootCmd.Flags().Bool("no-symbol", false, "Skip symbol pass")
	rootCmd.Flags().Bool("open", false, "Open browser")
	rootCmd.Flags().StringP("output", "o", "", "Write to file")
	rootCmd.Flags().String("padding-box", "", "Padding between box border and content")
	rootCmd.Flags().String("padding-root", "", "Padding around root content")
	rootCmd.Flags().Bool("tui", false, "Use terminal interface to explore the details")
	rootCmd.Flags().Bool("update-cache", false, "Update the cache file for the web UI")
	rootCmd.Flags().Bool("verbose", false, "Verbose output")
	rootCmd.Flags().Bool("version", false, "Show version")
	rootCmd.Flags().Bool("web", false, "use web interface to explore the details")
	rootCmd.Flags().String("width", "", "Width of the svg treemap")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json", "html", "svg"),
		"listen": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Suffix(":")
			default:
				return net.ActionPorts()
			}
		}),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
