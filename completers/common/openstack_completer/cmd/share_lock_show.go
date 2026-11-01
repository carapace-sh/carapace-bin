package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var share_lock_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details about a resource lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lock_showCmd).Standalone()

	share_lock_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_lock_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_lock_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_lock_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_lock_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_lock_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_lock_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_lock_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_lockCmd.AddCommand(share_lock_showCmd)

	carapace.Gen(share_lock_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns(share_lock_showCmd),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})

	carapace.Gen(share_lock_showCmd).PositionalCompletion(action.ActionShareLocks(share_lock_showCmd))
}
