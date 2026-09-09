package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swift",
	Short: "Swift compiler and package manager",
	Long:  "https://www.swift.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	common.AddCompilerFlags(rootCmd)
	common.CompilerFlagCompletion(rootCmd)

	rootCmd.Flags().BoolP("help", "h", false, "Show help information")
	rootCmd.Flags().Bool("version", false, "Show version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".swift"),
	)
}
