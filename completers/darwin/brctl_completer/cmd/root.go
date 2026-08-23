package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brctl",
	Short: "manage the CloudDocs daemon",
	Long:  "https://keith.github.io/xcode-manpages/brctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "List all non-local versions including those that are locally cached")
	rootCmd.Flags().Bool("collect-mobile-documents", false, "collect mobile documents")
	rootCmd.Flags().Bool("color", false, "turn on or off color use")
	rootCmd.Flags().String("database-path", "", "use the database at <db-path>")
	rootCmd.Flags().Bool("digest", false, "only print digest logs")
	rootCmd.Flags().String("filter", "", "only show lines matching predicate")
	rootCmd.Flags().String("home", "", "use this as the ~ prefix, to look for ~/L/")
	rootCmd.Flags().Bool("multiline", false, "turn on or off multiple line logging")
	rootCmd.Flags().String("name", "", "change the device name")
	rootCmd.Flags().String("output", "", "redirect output to <file-path>")
	rootCmd.Flags().Bool("page", false, "use paging")
	rootCmd.Flags().String("path", "", "use <logs-dir> instead of default")
	rootCmd.Flags().String("scope", "", "restrict the NSMDQ scope to DOCS, DATA, or BOTH")
	rootCmd.Flags().Bool("shorten", false, "shorten UUIDs, paths, etc")
	rootCmd.Flags().Bool("sysdiagnose", false, "do not collect what's already part of sysdiagnose")
	rootCmd.Flags().Bool("wait", false, "wait for new logs continuously")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"path":   carapace.ActionFiles(),
		"scope":  carapace.ActionValues("DOCS", "DATA", "BOTH"),
	})
}