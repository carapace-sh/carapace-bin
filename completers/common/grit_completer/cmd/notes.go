package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "List, show, add, or remove Git notes on commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notesCmd).Standalone()

	notesCmd.PersistentFlags().Bool("dry-run", false, "Preview `add` or `remove` without writing objects or refs")
	notesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	notesCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	notesCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	notesCmd.PersistentFlags().String("ref", "", "Notes ref (default: `core.notesRef` or `refs/notes/commits`)")
	rootCmd.AddCommand(notesCmd)
}
