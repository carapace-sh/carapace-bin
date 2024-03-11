package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyTagCmd = &cobra.Command{
	Use:   "verify-tag",
	Short: "Check the GPG signature of tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyTagCmd).Standalone()

	verifyTagCmd.Flags().String("format", "", "format to use for the output")
	verifyTagCmd.Flags().Bool("raw", false, "print raw gpg status output")
	verifyTagCmd.Flags().BoolP("verbose", "v", false, "print tag contents")
	rootCmd.AddCommand(verifyTagCmd)

	// TODO positional completion
}
