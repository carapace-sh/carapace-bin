package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "leaks",
	Short: "print information about leaked memory",
	Long:  "https://man.freebsd.org/cgi/man.cgi?leaks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("list", false, "Print the leaks as a list")
	rootCmd.Flags().Bool("groupByType", false, "Group children by type")
	rootCmd.Flags().Bool("nostacks", false, "Do not print backtraces")
	rootCmd.Flags().Bool("nosources", false, "Do not print sourceFile:lineNumber")
	rootCmd.Flags().Bool("quiet", false, "Do not print header")
	rootCmd.Flags().String("exclude", "", "Exclude symbol")
	rootCmd.Flags().String("outputGraph", "", "Path to output memory graph file")
	rootCmd.Flags().Bool("fullContent", false, "Include descriptions of content")
	rootCmd.Flags().Bool("noContent", false, "Do not print descriptions of content")
	rootCmd.Flags().Bool("fullStackHistory", false, "Full stack history")
	rootCmd.Flags().String("diffFrom", "", "Compare with memgraph file")
	rootCmd.Flags().String("traceTree", "", "Address to trace tree")
	rootCmd.Flags().Bool("referenceTree", false, "Reference tree")
	rootCmd.Flags().Bool("autoreleasePools", false, "Autorelease pools")
	rootCmd.Flags().String("debug", "", "Debug mode")
	rootCmd.Flags().Bool("conservative", false, "Conservative mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"outputGraph": carapace.ActionFiles(),
		"diffFrom":    carapace.ActionFiles(),
	})
}
