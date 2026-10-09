package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new share group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_createCmd).Standalone()

	share_group_createCmd.Flags().String("availability-zone", "", "Optional availability zone in which group should be created")
	share_group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_group_createCmd.Flags().String("description", "", "Share group description.")
	share_group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_group_createCmd.Flags().String("name", "", "Share group name")
	share_group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_group_createCmd.Flags().String("share-group-type", "", "Share group type name or ID of the share group to be created.")
	share_group_createCmd.Flags().String("share-network", "", "Specify share network name or id")
	share_group_createCmd.Flags().String("share-types", "", "Name or ID of share type(s).")
	share_group_createCmd.Flags().String("source-share-group-snapshot", "", "Share group snapshot name or ID to create the share group from.")
	share_group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_group_createCmd.Flags().Bool("wait", false, "Wait for share group creation")
	share_groupCmd.AddCommand(share_group_createCmd)
}
