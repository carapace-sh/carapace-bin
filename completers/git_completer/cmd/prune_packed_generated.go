package cmd

import (
	"github.com/spf13/cobra"
)

var prune_packedCmd = &cobra.Command{
	Use:   "prune-packed",
	Short: "Remove extra objects that are already in pack files",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	prune_packedCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	prune_packedCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	rootCmd.AddCommand(prune_packedCmd)
}
