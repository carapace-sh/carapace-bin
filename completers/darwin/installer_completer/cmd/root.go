package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "installer",
	Short: "system software and package installer tool",
	Long:  "https://keith.github.io/xcode-manpages/installer.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allow", false, "Allow")
	rootCmd.Flags().String("applyChoiceChangesXML", "", "Path to XML file")
	rootCmd.Flags().Bool("config", false, "Display configuration")
	rootCmd.Flags().Bool("dominfo", false, "Display domain information")
	rootCmd.Flags().Bool("dumplog", false, "Dump log")
	rootCmd.Flags().String("file", "", "Path to file")
	rootCmd.Flags().Bool("help", false, "Display help")
	rootCmd.Flags().String("lang", "", "ISO language code")
	rootCmd.Flags().Bool("listiso", false, "List ISO languages")
	rootCmd.Flags().String("pkg", "", "Path to package")
	rootCmd.Flags().Bool("pkginfo", false, "Display package information")
	rootCmd.Flags().Bool("plist", false, "Display as plist")
	rootCmd.Flags().String("query", "", "Query flag")
	rootCmd.Flags().String("showChoicesAfterApplyingChangesXML", "", "Path to XML file")
	rootCmd.Flags().Bool("showChoicesXML", false, "Display choices XML")
	rootCmd.Flags().String("target", "", "Target device")
	rootCmd.Flags().Bool("verbose", false, "Verbose")
	rootCmd.Flags().Bool("verboseR", false, "Verbose (regular)")
	rootCmd.Flags().Bool("vers", false, "Display version")
	rootCmd.Flags().Bool("volinfo", false, "Display volume information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"applyChoiceChangesXML":              carapace.ActionFiles(),
		"file":                               carapace.ActionFiles(),
		"lang":                               os.ActionLanguages(),
		"pkg":                                carapace.ActionFiles(),
		"showChoicesAfterApplyingChangesXML": carapace.ActionFiles(),
	})
}
