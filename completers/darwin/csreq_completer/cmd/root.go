package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csreq",
	Short: "expert tool for manipulating Code Signing Requirement data",
	Long:  "https://keith.github.io/xcode-manpages/csreq.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("binary", "b", "", "Write the requirement in binary form to the path")
	rootCmd.Flags().StringP("requirement", "r", "", "Specify the input requirement")
	rootCmd.Flags().BoolP("text", "t", false, "Write the requirement as text to stdout")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity of output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"binary": carapace.ActionFiles(),
		"r":      carapace.ActionFiles(),
	})
}
