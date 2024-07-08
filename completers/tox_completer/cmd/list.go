package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l"},
	Short: "list environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	add_common_flags(listCmd)

	// TODO: multiple commands have m, f, and skip-env ; maybe consolidate to another common function?
	listCmd.Flags().StringArrayS("m", "m", []string{}, "labels to evaluate (default: [])")
	listCmd.Flags().StringArrayS("f", "f", []string{}, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	listCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")
	listCmd.Flags().BoolS("d", "d", false, "list just default envs (default: False)")

	rootCmd.AddCommand(listCmd)
}
