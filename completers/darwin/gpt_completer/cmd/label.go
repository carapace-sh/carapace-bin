package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "label a partition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelCmd).Standalone()
	rootCmd.AddCommand(labelCmd)

	labelCmd.Flags().Bool("a", false, "Apply to all matching")
	labelCmd.Flags().StringP("file", "f", "", "Read label from file")
	labelCmd.Flags().StringP("label", "l", "", "Label string")

	carapace.Gen(labelCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}