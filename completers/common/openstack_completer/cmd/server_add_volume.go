package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Add volume to server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_volumeCmd).Standalone()

	server_add_volumeCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_add_volumeCmd.Flags().String("device", "", "Server internal device name for volume")
	server_add_volumeCmd.Flags().Bool("disable-delete-on-termination", false, "Do not delete the volume when the server is destroyed (supported by --os-compute-api-version 2.79 or above)")
	server_add_volumeCmd.Flags().Bool("enable-delete-on-termination", false, "Delete the volume when the server is destroyed (supported by --os-compute-api-version 2.79 or above)")
	server_add_volumeCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_add_volumeCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_add_volumeCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_add_volumeCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_add_volumeCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_add_volumeCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_add_volumeCmd.Flags().String("tag", "", "Tag for the attached volume (supported by --os-compute-api-version 2.49 or above)")
	server_add_volumeCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_addCmd.AddCommand(server_add_volumeCmd)
}
