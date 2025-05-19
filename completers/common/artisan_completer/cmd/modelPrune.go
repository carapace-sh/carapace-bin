package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modelPruneCmd = &cobra.Command{
	Use:   "model:prune [--model [MODEL]] [--except [EXCEPT]] [--path [PATH]] [--chunk [CHUNK]] [--pretend]",
	Short: "Prune models that are no longer needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modelPruneCmd).Standalone()

	modelPruneCmd.Flags().String("chunk", "", "The number of models to retrieve per chunk of models to be deleted")
	modelPruneCmd.Flags().String("except", "", "Class names of the models to be excluded from pruning")
	modelPruneCmd.Flags().String("model", "", "Class names of the models to be pruned")
	modelPruneCmd.Flags().String("path", "", "Absolute path(s) to directories where models are located")
	modelPruneCmd.Flags().Bool("pretend", false, "Display the number of prunable records found instead of deleting them")
	rootCmd.AddCommand(modelPruneCmd)
}
