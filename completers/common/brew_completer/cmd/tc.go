package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tcCmd = &cobra.Command{
	Use:   "tc",
	Short: "Check for typechecking errors using Sorbet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcCmd).Standalone()

	tcCmd.Flags().Bool("debug", false, "Display any debugging information.")
	tcCmd.Flags().String("dir", "", "Typecheck all files in a specific directory.")
	tcCmd.Flags().String("file", "", "Typecheck a single file.")
	tcCmd.Flags().Bool("fix", false, "Automatically fix type errors.")
	tcCmd.Flags().Bool("help", false, "Show this message.")
	tcCmd.Flags().String("ignore", "", "Ignores input files that contain the given string in their paths (relative to the input path passed to Sorbet).")
	tcCmd.Flags().Bool("quiet", false, "Silence all non-critical errors.")
	tcCmd.Flags().String("suggest-typed", "", "Try upgrading `typed` sigils.")
	tcCmd.Flags().String("update", "", "Update RBI files.")
	tcCmd.Flags().Bool("update-all", false, "Update all RBI files rather than just updated gems.")
	tcCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(tcCmd)
}
