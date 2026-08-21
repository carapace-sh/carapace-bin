package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_member_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show a particular project associated with image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_member_getCmd).Standalone()

	image_member_getCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_member_getCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_member_getCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_member_getCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_member_getCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_member_getCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_member_getCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_member_getCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	image_member_getCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_memberCmd.AddCommand(image_member_getCmd)
}
