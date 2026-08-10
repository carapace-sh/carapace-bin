package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_dumpSymbolGraphCmd = &cobra.Command{
	Use:   "dump-symbol-graph",
	Short: "Dump symbol graphs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_dumpSymbolGraphCmd).Standalone()
	package_dumpSymbolGraphCmd.Flags().SetInterspersed(false)

	package_dumpSymbolGraphCmd.Flags().String("extension-block-symbol-behavior", "", "Emit extension block symbols or omit them")
	package_dumpSymbolGraphCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_dumpSymbolGraphCmd.Flags().Bool("include-spi-symbols", false, "Add symbols with SPI information to the symbol graph")
	package_dumpSymbolGraphCmd.Flags().String("minimum-access-level", "", "Include symbols with this access level or more")
	package_dumpSymbolGraphCmd.Flags().String("output-dir", "", "Specify symbol graph output directory")
	package_dumpSymbolGraphCmd.Flags().Bool("pretty-print", false, "Pretty-print the output JSON")
	package_dumpSymbolGraphCmd.Flags().Bool("skip-inherited-docs", false, "Skip emitting doc comments for inherited members")
	package_dumpSymbolGraphCmd.Flags().Bool("skip-synthesized-members", false, "Skip members inherited through classes or default implementations")
	package_dumpSymbolGraphCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_dumpSymbolGraphCmd)

	carapace.Gen(package_dumpSymbolGraphCmd).FlagCompletion(carapace.ActionMap{
		"extension-block-symbol-behavior": carapace.ActionValues("emit-extension-block-symbols", "omit-extension-block-symbols"),
		"minimum-access-level":            carapace.ActionValues("private", "fileprivate", "internal", "package", "public", "open"),
		"output-dir":                      carapace.ActionDirectories(),
	})
}
