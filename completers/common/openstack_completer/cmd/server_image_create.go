package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_image_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new server disk image from an existing server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_image_createCmd).Standalone()

	server_image_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_image_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_image_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_image_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_image_createCmd.Flags().String("name", "", "Name of new disk image (default: server name)")
	server_image_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_image_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_image_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_image_createCmd.Flags().String("property", "", "Set a new property to meta_data.json on the metadata server (repeat option to set multiple values)")
	server_image_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_image_createCmd.Flags().Bool("wait", false, "Wait for operation to complete")
	server_imageCmd.AddCommand(server_image_createCmd)
}
