package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagCmd = &cobra.Command{
	Use:   "diag",
	Short: "Create diagnostic report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagCmd).Standalone()

	diagCmd.Flags().BoolP("append", "a", false, "Use recipes specified with -r in combination with built-in recipes")
	diagCmd.Flags().BoolP("check", "c", false, "Validate recipes")
	diagCmd.Flags().BoolP("debug", "d", false, "Print debug information")
	diagCmd.Flags().BoolP("force", "f", false, "Delete target file or directory if it already exists")
	diagCmd.Flags().BoolP("list", "l", false, "Show information about available recipes")
	diagCmd.Flags().Bool("log-prefix", false, "Add a severity character at the beginning of log messages")
	diagCmd.Flags().Bool("no-color", false, "Do not emit ANSI color codes")
	diagCmd.Flags().StringP("output", "o", "", "Diagnostic report directory or zip file archive path")
	diagCmd.Flags().String("purge", "", "Use a custom purge file to remove sensitive information")
	diagCmd.Flags().StringArrayP("recipe", "r", nil, "Recipe to use")
	diagCmd.Flags().StringArrayP("tag", "t", nil, "Consider only recipes containing TAG")
	diagCmd.Flags().StringP("zip", "z", "", "Create zip archive for diagnostic report in PATH")
	rootCmd.AddCommand(diagCmd)

	carapace.Gen(diagCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"purge":  carapace.ActionFiles(),
		"recipe": carapace.ActionFiles(),
		"zip":    carapace.ActionFiles(),
	})
}
