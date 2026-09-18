package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_type_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new share group type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_type_createCmd).Standalone()

	share_group_type_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_group_type_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_group_type_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_group_type_createCmd.Flags().String("group-specs", "", "Share Group type extra specs by key and value.")
	share_group_type_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_group_type_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_group_type_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_group_type_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_group_type_createCmd.Flags().String("public", "", "Make type accessible to the public (default true).")
	share_group_type_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_group_typeCmd.AddCommand(share_group_type_createCmd)
}
