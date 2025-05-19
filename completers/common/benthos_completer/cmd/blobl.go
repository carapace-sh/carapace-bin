package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bloblCmd = &cobra.Command{
	Use:   "blobl",
	Short: "Execute a Bloblang mapping on documents consumed via stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bloblCmd).Standalone()

	bloblCmd.Flags().StringP("file", "f", "", "execute a mapping from a file.")
	bloblCmd.Flags().String("max-token-length", "", "Set the buffer size for document lines.")
	bloblCmd.Flags().BoolP("pretty", "p", false, "pretty-print output.")
	bloblCmd.Flags().BoolP("raw", "r", false, "consume raw strings.")
	bloblCmd.Flags().StringP("threads", "t", "", "the number of processing threads to use, when >1 ordering is no longer guaranteed.")
	rootCmd.AddCommand(bloblCmd)

	carapace.Gen(bloblCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
