package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve a conflicted file with an external merge tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.Flags().BoolP("list", "l", false, "Instead of resolving one conflict, list all the conflicts")
	resolveCmd.Flags().BoolP("quiet", "q", false, "Do not print the list of remaining conflicts (if any) after resolving a conflict")
	resolveCmd.Flags().StringP("revision", "r", "", "")
	rootCmd.AddCommand(resolveCmd)
}
