package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var evalCmd = &cobra.Command{
	Use:     "eval",
	Short:   "evaluate a Nix expression",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evalCmd).Standalone()

	evalCmd.Flags().String("apply", "", "Apply the function expr to each argument")
	evalCmd.Flags().Bool("json", false, "Produce output in JSON format, suitable for consumption by another program")
	evalCmd.Flags().Bool("raw", false, "Print strings without quotes or escaping")
	evalCmd.Flags().Bool("read-only", false, "Do not instantiate each evaluated derivation")
	evalCmd.Flags().String("write-to", "", "Write a string or attrset of strings to path")
	rootCmd.AddCommand(evalCmd)

	addEvaluationFlags(evalCmd)
	addFlakeFlags(evalCmd)
	addLoggingFlags(evalCmd)

	// TODO flag completion

	// TODO positional completion
}
