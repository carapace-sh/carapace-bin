package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_object_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a metadef object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_object_createCmd).Standalone()

	image_metadef_object_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_metadef_object_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_metadef_object_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_metadef_object_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_metadef_object_createCmd.Flags().String("namespace", "", "Metadef namespace to create the object in (name)")
	image_metadef_object_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_metadef_object_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_metadef_object_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_metadef_object_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_metadef_objectCmd.AddCommand(image_metadef_object_createCmd)
}
