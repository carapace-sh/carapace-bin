package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reflog_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete single entries from the reflog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_deleteCmd).Standalone()

	reflog_deleteCmd.Flags().BoolP("dry-run", "n", false, "Do not actually prune any entries")
	reflog_deleteCmd.Flags().Bool("rewrite", false, "Adjust a reflogs \"old\" SHA-1 to be equal to the \"new\" SHA-1 field of the entry that now precedes it")
	reflog_deleteCmd.Flags().Bool("updateref", false, "Update the reference to the value of the top reflog entry")
	reflog_deleteCmd.Flags().Bool("verbose", false, "Print extra information on screen.")

	reflogCmd.AddCommand(reflog_deleteCmd)
}
