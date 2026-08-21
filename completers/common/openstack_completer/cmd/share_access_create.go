package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_access_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new share access rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_access_createCmd).Standalone()

	share_access_createCmd.Flags().String("access-level", "", "Share access level (\"rw\" and \"ro\" access levels are supported).")
	share_access_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_access_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_access_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_access_createCmd.Flags().Bool("lock-deletion", false, "When enabled, a 'delete' lock will be placed against the rule and the rule cannot be deleted while the lock exists.")
	share_access_createCmd.Flags().String("lock-reason", "", "Reason for locking the access rule.")
	share_access_createCmd.Flags().Bool("lock-visibility", false, "Whether the sensitive fields of the access rule redacted to other users.")
	share_access_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_access_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_access_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_access_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_access_createCmd.Flags().String("properties", "", "Space separated list of key=value pairs of properties.")
	share_access_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_access_createCmd.Flags().Bool("wait", false, "Wait for share access rule creation.")
	share_accessCmd.AddCommand(share_access_createCmd)
}
