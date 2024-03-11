package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var typecheckCmd = &cobra.Command{
	Use:     "typecheck",
	Short:   "Check for typechecking errors using Sorbet",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(typecheckCmd).Standalone()

	typecheckCmd.Flags().Bool("debug", false, "Display any debugging information.")
	typecheckCmd.Flags().Bool("dir", false, "Typecheck all files in a specific directory.")
	typecheckCmd.Flags().Bool("file", false, "Typecheck a single file.")
	typecheckCmd.Flags().Bool("fix", false, "Automatically fix type errors.")
	typecheckCmd.Flags().Bool("help", false, "Show this message.")
	typecheckCmd.Flags().Bool("ignore", false, "Ignores input files that contain the given string in their paths (relative to the input path passed to Sorbet).")
	typecheckCmd.Flags().Bool("quiet", false, "Silence all non-critical errors.")
	typecheckCmd.Flags().Bool("suggest-typed", false, "Try upgrading `typed` sigils.")
	typecheckCmd.Flags().Bool("update", false, "Update RBI files.")
	typecheckCmd.Flags().Bool("update-all", false, "Update all RBI files rather than just updated gems.")
	typecheckCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(typecheckCmd)
}
