package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_lock_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Lock a resource action from occurring on a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lock_createCmd).Standalone()

	share_lock_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_lock_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_lock_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_lock_createCmd.Flags().String("lock-reason", "", "Reason for the resource lock.")
	share_lock_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_lock_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_lock_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_lock_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_lock_createCmd.Flags().String("resource-action", "", "Action to lock on the resource (default=\"delete\")")
	share_lock_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_lockCmd.AddCommand(share_lock_createCmd)
}
