package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Helper to hash installer files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hashCmd).Standalone()

	hashCmd.Flags().StringP("file", "f", "", "File to be hashed")
	hashCmd.Flags().BoolP("msix", "m", false, "Input file will be treated as msix; signature hash will be provided if signed")
	rootCmd.AddCommand(hashCmd)

	carapace.Gen(hashCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
