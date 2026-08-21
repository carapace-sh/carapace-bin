package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_quota_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show Quota",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_quota_showCmd).Standalone()

	share_quota_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_quota_showCmd.Flags().Bool("defaults", false, "Show the default quotas for the project.")
	share_quota_showCmd.Flags().Bool("detail", false, "Optional flag to indicate whether to show quota in detail.")
	share_quota_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_quota_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_quota_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_quota_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_quota_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_quota_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_quota_showCmd.Flags().String("share-type", "", "Name or ID of a share type to list the quotas for.")
	share_quota_showCmd.Flags().String("user", "", "Name or ID of user to list the quotas for.")
	share_quota_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_quotaCmd.AddCommand(share_quota_showCmd)
}
