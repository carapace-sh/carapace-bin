package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clangHtmlReportCmd = &cobra.Command{
	Use:   "clang-html-report",
	Short: "Generate html report to html_report folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clangHtmlReportCmd).Standalone()

	rootCmd.AddCommand(clangHtmlReportCmd)
}
