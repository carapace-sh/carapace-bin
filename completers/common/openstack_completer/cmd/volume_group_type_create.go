package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_type_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a volume group type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_type_createCmd).Standalone()

	volume_group_type_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_group_type_createCmd.Flags().String("description", "", "Description of the volume group type.")
	volume_group_type_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_group_type_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_group_type_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_group_type_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_group_type_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_group_type_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_group_type_createCmd.Flags().Bool("private", false, "Volume group type is not available to other projects")
	volume_group_type_createCmd.Flags().Bool("public", false, "Volume group type is available to other projects (default)")
	volume_group_type_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_group_typeCmd.AddCommand(volume_group_type_createCmd)
}
